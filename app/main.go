package main

import (
	"github.com/go-martini/martini"
	"net/http"
	"appengine"
	"fmt"
	"github.com/bearchinc/datastore-model"
	"github.com/heckfer/ProjectCars/model"
	"log"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/binding"
)

func init() {

	martini := martini.Classic()
	martini.Use(render.Renderer(render.Options{IndentJSON: true, }))

	martini.Post("/create_post", binding.Bind(model.Post{}), CreatePostHandler)
	martini.Get("/list_posts", ListPostsHandler)

	http.Handle("/", martini)
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