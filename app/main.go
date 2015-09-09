package main

import (
	"net/http"
	"github.com/heckfer/fala-com-meu-carro/model"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/binding"
	"github.com/go-martini/martini"
	"github.com/heckfer/fala-com-meu-carro/handler"
	"github.com/heckfer/fala-com-meu-carro/middleware"
)

func init() {

	m := martini.Classic()
	m.Use(render.Renderer(render.Options{IndentJSON: true, }))
	m.Use(middleware.AppengineContextProvider)
	m.Use(middleware.AppxProvider)

	m.Group("/posts", func(r martini.Router) {
		r.Post("/create", binding.Bind(model.Post{}), handler.CreatePostHandler)
		r.Get("/list", handler.ListPostsHandler)
		r.Get("/list/:plate", handler.ListPostsByCarPlateHandler)
		r.Put("/:post_id", handler.FlagPostHandler)
	})

	http.Handle("/", m)
}
