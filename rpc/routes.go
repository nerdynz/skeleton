package main

import (
	"embed"
	"errors"
	"net/http"

	"github.com/go-zoo/bone"
	"github.com/nerdynz/datastore"
	"github.com/nerdynz/flow"
	"github.com/nerdynz/router"
	"github.com/nerdynz/security"
	"github.com/nerdynz/skeleton/rpc/access"

	"github.com/unrolled/render"
)

//go:embed templates
var templates embed.FS

func Routes(store *datastore.Datastore) *bone.Mux {
	// func Routes(store *datastore.Datastore, w *wclient.Client) *bone.Mux {
	key := &Key{
		Store: store,
	}

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
	// personTwirp := person.NewServer(store)

	// The generated code includes a method, PathPrefix(), which
	// can be used to mount your service on a mux.
	r.Mux.Handle(accessTwirp.PathPrefix(), accessTwirp)

	r.GET("/robots.txt", robots, security.NoAuth)

	r.GET("/:a/:b/:c/:d/:e/:f/:g/:h", spa, security.NoAuth)
	r.GET("/:a/:b/:c/:d/:e/:f/:g", spa, security.NoAuth)
	r.GET("/:a/:b/:c/:d/:e/:f", spa, security.NoAuth)
	r.GET("/:a/:b/:c/:d/:e", spa, security.NoAuth)
	r.GET("/:a/:b/:c/:d", spa, security.NoAuth)
	r.GET("/:a/:b/:c", spa, security.NoAuth)
	r.GET("/:a/:b", spa, security.NoAuth)
	r.GET("/:a", spa, security.NoAuth)
	r.GET("/", spa, security.NoAuth)

	// IMAGE HANDLING
	r.POST("/upload/image", uploadImage, security.NoAuth)

	r.Mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

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
