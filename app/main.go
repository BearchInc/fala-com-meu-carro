package main

import (
	"net/http"
	"appengine"
	"fmt"
	"log"
	"github.com/bearchinc/datastore-model"
	"github.com/heckfer/ProjectCars/model"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/binding"
	"github.com/go-martini/martini"
)

func init() {

	m := martini.Classic()
	m.Use(render.Renderer(render.Options{IndentJSON: true, }))

	m.Group("/posts", func(r martini.Router) {
		r.Post("/create", binding.Bind(model.Post{}), CreatePostHandler)
		r.Get("/list", ListPostsHandler)
		r.Get("/list/:plate", ListPostsByCarPlateHandler)
	})

	http.Handle("/", m)
}


func CreatePostHandler(c martini.Context, req *http.Request, r render.Render, post model.Post) {

	context := getAppengineContext(req)

	err := db.NewDatastore(context).Create(&post)
	if err != nil {
		log.Printf("Error: %+v", err)
		r.JSON(500, "Error")
	} else {
		r.JSON(200, post)
	}
}

func ListPostsByCarPlateHandler(c martini.Context, req *http.Request, r render.Render, params martini.Params) {

	context := getAppengineContext(req)

	posts := model.Posts{}
	carPlate := params["plate"]
	err := db.NewDatastore(context).Query(db.From(new(model.Post)).Filter("CarPlate=", carPlate)).All(&posts)

	if err != nil {
		log.Printf("Error: %+v", err)
		r.JSON(500, "Error")
	} else {
		r.JSON(200, posts)
	}
}


func ListPostsHandler(c martini.Context, req *http.Request, r render.Render) {

	context := getAppengineContext(req)

	posts := model.Posts{}
	err := db.NewDatastore(context).Query(db.From(new(model.Post))).All(&posts)

	if err != nil {
		log.Printf("Error: %+v", err)
		r.JSON(500, "Error")
	} else {
		r.JSON(200, posts)
	}
}



func getAppengineContext(request *http.Request) appengine.Context {
	gae := appengine.NewContext(request)
	namespace := appengine.ModuleName(gae)

	context, err := appengine.Namespace(gae, namespace)

	if err != nil {
		panic(fmt.Sprintf("Could not create GAE context: %v", err))
	}

	return context
}