package main

import (
	"embed"
	"errors"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-zoo/bone"
	"github.com/nerdynz/datastore"
	"github.com/nerdynz/flow"
	"github.com/nerdynz/router"
	"github.com/nerdynz/security"
	"github.com/nerdynz/skeleton/rpc/access"
	"github.com/nerdynz/skeleton/rpc/page"
	"github.com/nerdynz/skeleton/rpc/ssr"

	"github.com/unrolled/render"
)

//go:embed templates
var templates embed.FS

//go:embed all:dist
var dist embed.FS

func Routes(store *datastore.Datastore) *bone.Mux {

	// func Routes(store *datastore.Datastore, w *wclient.Client) *bone.Mux {
	key := &Key{
		Store: store,
	}

	// fsysStatic, _ := fs.Sub(static, "static")
	// fsysFrontend, err := fs.Sub(dist, "frontend/client")

	// fsysServer, _ := fs.Sub(dist, "frontend/server")

	var renderer = render.New(render.Options{
		Layout:     "application",
		FileSystem: render.FS(templates),
		Extensions: []string{".html"},
		// Funcs: []template.FuncMap{
		// 	HelperFuncs,
		// },
		// prevent having to rebuild for every template reload... This is an important setting for development speed
		IsDevelopment:               store.Settings.IsDevelopment(),
		RequirePartials:             store.Settings.IsDevelopment(),
		RequireBlocks:               store.Settings.IsDevelopment(),
		RenderPartialsWithoutPrefix: true,
	})

	r := router.New(renderer, store, key, true)
	// Scaffold routes

	accessTwirp := access.NewServer(store, key)
	pageTwirp := page.NewRpcServer(store)
	// personTwirp := person.NewServer(store)

	// The generated code includes a method, PathPrefix(), which
	// can be used to mount your service on a mux.
	r.Mux.Handle(accessTwirp.PathPrefix(), accessTwirp)
	r.Mux.Handle(pageTwirp.PathPrefix(), pageTwirp)

	ssr, err := ssr.NewRenderer(dist)
	if err != nil {
		slog.Error("failed to setup ssr renderer", "err", err)
		os.Exit(1)
	}

	r.GET("/robots.txt", robots, security.NoAuth)

	r.GET("/metadata", ssr.Metadata, security.NoAuth)
	r.GET("/page/:slug", ssr.JSON, security.NoAuth)
	r.GET("/:slug", ssr.Handle, security.NoAuth)
	r.GET("/", ssr.Handle, security.NoAuth)

	// IMAGE HANDLING
	r.POST("/upload/image", uploadImage, security.NoAuth)

	r.Mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	return r.Mux
}

func spa(w http.ResponseWriter, req *http.Request, ctx *flow.Flow, store *datastore.Datastore) {
	ctx.HTML("spa", 200)
}

func testErrorFunc() error {
	err := errors.New("generate a test error")
	return err
}

func ErrorTest(w http.ResponseWriter, req *http.Request, ctx *flow.Flow, store *datastore.Datastore) {
	err := testErrorFunc()
	if err != nil {
		ctx.ErrorJSON(http.StatusInternalServerError, "Test error is working", err)
		return
	}

	ctx.JSON(200, "Working")
}

func robots(w http.ResponseWriter, req *http.Request, ctx *flow.Flow, store *datastore.Datastore) {
	robotsTxt := `
	User-agent: *
	Disallow: /
	`
	ctx.Renderer.Text(w, 200, robotsTxt)
}
