package main

import (
	"github.com/go-martini/martini"
	"net/http"
	"appengine"
	"fmt"
	"github.com/bearchinc/datastore-model"
	"github.com/heckfer/ProjectCars/model"
)

func init() {

	router := martini.Classic()

	router.Post("/create_post", CreatePostHandler)
	router.Get("/list_posts", ListPostsHandler)

	http.Handle("/", router)
}


func CreatePostHandler(c martini.Context, req *http.Request) string {
	gae := appengine.NewContext(req)
	namespace := appengine.ModuleName(gae)
	context, err := appengine.Namespace(gae, namespace)
	if err != nil {
		panic(fmt.Sprintf("Could not create GAE context: %v", err))
	}

	post := &model.Post{
		CarPlate: "IWI-5585",
		Message: "Lisardo perdeu outro cartão!",
	}

	db.NewDatastore(context).Create(post)

	return "OK"
}

func ListPostsHandler(c martini.Context, req *http.Request) string {
	gae := appengine.NewContext(req)
	namespace := appengine.ModuleName(gae)
	context, err := appengine.Namespace(gae, namespace)
	if err != nil {
		panic(fmt.Sprintf("Could not create GAE context: %v", err))
	}

	post := &model.Post{
		CarPlate: "IWI-5585",
		Message: "Lisardo perdeu outro cartão!",
	}

	db.NewDatastore(context).Create(post)

	return "OK"
}