package handler
import (
	"github.com/go-martini/martini"
	"net/http"
	"github.com/martini-contrib/render"
	"github.com/heckfer/fala-com-meu-carro/model"
	"github.com/bearchinc/datastore-model"
	"log"
	"appengine"
	"fmt"
	"strings"
)

func CreatePostHandler(c martini.Context, req *http.Request, r render.Render, post model.Post) {
	context := getAppengineContext(req)

	post.CarPlate = strings.ToUpper(post.CarPlate)

	err := db.NewDatastore(context).Create(&post)
	if err != nil {
		log.Printf("Error: %+v", err)
		r.JSON(500, "Error")
	} else {
		r.JSON(200, post)
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