package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	// nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/lmittmann/tint"
	"github.com/r3labs/sse"
	"github.com/sirupsen/logrus"

	"golang.org/x/exp/slices"

	// "github.com/nerdynz/skeleton/rpc/block"
	"github.com/nerdynz/datastore"
	"github.com/nerdynz/rcache"
	"github.com/nerdynz/security"
	"github.com/nerdynz/trove"

	// redisv9 "github.com/redis/go-redis/v9"
	"github.com/twitchtv/twirp"
	"github.com/urfave/negroni"
)

var settings *trove.Settings
var store *datastore.Datastore

func main() {
	settings = trove.Load()
	// create a new logger
	// logger := slog.New(tint.NewHandler(w, nil))

	// set global logger with custom options
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level:      slog.LevelDebug,
			TimeFormat: time.Kitchen,
			NoColor:    !settings.IsDevelopment(),
		}),
	))

	nL := &negroni.Logger{
		ALogger: &logSlog{
			slog.Default(),
		},
	}
	nL.SetFormat(negroni.LoggerDefaultFormat)
	n := negroni.New(negroni.NewRecovery(), nL)

	// tf := time.Stamp
	// if settings.IsDevelopment() {
	// 	logger.Info("\n\n===================\n\n** SERVER IS DEV ** \n\n===================\n\n")
	// 	tf = time.TimeOnly
	// }

	// logger.SetFormatter(&nested.Formatter{
	// 	HideKeys:        true,
	// 	TimestampFormat: tf,
	// })

	// filestorage := fileupload.NewLocalFileStorage(settings.Get("ATTACHMENTS_FOLDER"), "/attachments/")
	cache := rcache.New(settings.Get("CACHE_URL"), logrus.StandardLogger())

	store = datastore.New(settings, cache, nil)
	sseSrv := NewSseServer()
	store.Publisher = sseSrv
	defer sseSrv.Close()

	// rclient := redisv9.NewUniversalClient(&redisv9.UniversalOptions{
	// 	Addrs:        []string{"localhost:6379"},
	// 	Username:     "",
	// 	Password:     "RedisPassw0rd",
	// 	DB:           0,
	// 	WriteTimeout: time.Second * 30,
	// 	ReadTimeout:  time.Second * 30,
	// })

	// b, err := redis.NewRedisBackend(rclient, redis.WithAutoExpiration(time.Minute))
	// if err != nil {
	// 	panic(err)
	// }

	// Run worker
	// w := RunWorker(b)

	// ctx, cancel := context.WithCancel(context.Background())
	// go w.Start(ctx)

	// routes := Routes(store, client.New(b))

	routes := Routes(store)
	routes.HandleFunc("/events", sseSrv.sseServer.HTTPHandler)

	// Start diagnostic server under /diag
	// routes.Handle("/diag/", http.StripPrefix("/diag", diag.NewServeMux(b)))
	n.UseHandler(routes)
	http.ListenAndServe(":"+settings.GetWithDefault("DEVELOPMENT_PORT", "8080"), n)

	// cancel()
	// if err := w.WaitForCompletion(); err != nil {
	// 	logrus.Info("could not stop worker", err)
	// }
}

// func MyMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
// 	// do some stuff before
// 	next(rw, r)
// 	// do some stuff after
// }

var bypassUrls = []string{"/twirp/nerdynz.Access/Login", "/twirp/nerdynz.Access/ValidSites"}

func WithAuthorization(base http.Handler, key security.Key) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if slices.Contains(bypassUrls, r.URL.Path) {
			base.ServeHTTP(w, r)
			return
		}

		ctx := r.Context()
		auth := r.Header.Get("Authorization")

		padlock := security.New(r, settings, key)
		if !padlock.IsLoggedIn() {
			twirp.WriteError(w, twirp.NewError(twirp.Unauthenticated, "not logged in"))
			return
		}

		personUlid, err := padlock.LoggedInUserULID()
		if err != nil {
			twirp.WriteError(w, twirp.NewError(twirp.Unauthenticated, "not logged in"))
			return
		}
		ctx = context.WithValue(ctx, "person_ulid", personUlid)

		ctx = context.WithValue(ctx, "authorization", auth)
		siteUlid, err := padlock.SiteULID()
		if err != nil {
			twirp.WriteError(w, twirp.NewError(twirp.Unauthenticated, err.Error()))
			return
		}
		ctx = context.WithValue(ctx, "site_ulid", siteUlid)

		r = r.WithContext(ctx)
		base.ServeHTTP(w, r)
	})
}

type SseServer struct {
	sseServer *sse.Server
}

func NewSseServer() *SseServer {
	sseSrv := sse.New()
	sseSrv.AutoReplay = false
	// sseSrv.EncodeBase64 = true
	sseSrv.CreateStream("messages")

	return &SseServer{
		sseServer: sseSrv,
	}
}

func (ssesrv *SseServer) Publish(siteUlid string, entity string, messageType string, ids []string) error {
	b, err := json.Marshal(struct {
		Entity      string   `json:"entity"`
		MessageType string   `json:"messageType"`
		Ids         []string `json:"ids"`
	}{
		entity,
		messageType,
		ids,
	})
	if err != nil {
		return err
	}
	event := &sse.Event{
		Data: b,
	}
	ssesrv.sseServer.Publish("messages", event)
	return nil
}

func (sse *SseServer) Close() {
	sse.sseServer.Close()
}

type logSlog struct {
	logger *slog.Logger
}

func (ls *logSlog) Printf(format string, v ...any) {
	ls.logger.Info(format)
}

func (ls *logSlog) Println(v ...any) {
	output, ok := v[0].(string)
	if ok {
		logLevel := slog.LevelInfo

		// split into something useful
		spl := strings.Split(output, " | ")
		statusStr := spl[1]
		if status, err := strconv.Atoi(statusStr); err == nil {
			if status >= 400 {
				logLevel = slog.LevelError
			}
		}

		duration := strings.TrimPrefix(spl[2], "\t ")
		endpoint := strings.Split(spl[4], " ")
		// ls.logger.Info(output)
		ls.logger.Log(context.Background(), logLevel, "RPC", "duration", duration, "status", statusStr, "verb", endpoint[0], "endpoint", endpoint[1])
	} else {
		ls.logger.Debug("NOT IMPLEMENTED PRINT LINE NEGRONI")
	}
}
