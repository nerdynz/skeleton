package main

import (
	"context"
	"net/http"

	"github.com/nerdynz/skeleton/rpc/access"
	"github.com/nerdynz/skeleton/rpc/person"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/slices"

	// "github.com/nerdynz/skeleton/rpc/block"
	"github.com/nerdynz/datastore"
	"github.com/nerdynz/rcache"
	"github.com/nerdynz/security"
	"github.com/nerdynz/trove"
	"github.com/twitchtv/twirp"
	"github.com/urfave/negroni"
)

var settings *trove.Settings
var store *datastore.Datastore

func main() {
	n := negroni.Classic()
	settings = trove.Load()
	logger := logrus.New()

	// filestorage := fileupload.NewLocalFileStorage(settings.Get("ATTACHMENTS_FOLDER"), "/attachments/")
	cache := rcache.New(settings.Get("CACHE_URL"), logger)
	store = datastore.New(logger, settings, cache, nil)
	store.TurnOnLogging()

	key := &Key{
		Store: store,
	}

	accessTwirp := access.NewServer(store, key)
	personTwirp := person.NewServer(store)
	// autocompleteTwirp := autocomplete.NewServer(store)
	// blockTwirp := block.NewServer(store)

	mux := http.NewServeMux()
	// The generated code includes a method, PathPrefix(), which
	// can be used to mount your service on a mux.
	mux.Handle(accessTwirp.PathPrefix(), WithAuthorization(accessTwirp, key))
	mux.Handle(personTwirp.PathPrefix(), WithAuthorization(personTwirp, key))
	// mux.Handle(autocompleteTwirp.PathPrefix(), WithAuthorization(autocompleteTwirp, key))
	n.UseHandler(mux)

	n.Run(":4500")
}

func MyMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// do some stuff before
	next(rw, r)
	// do some stuff after
}

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
