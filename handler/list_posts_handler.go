package handler
import (
	"github.com/go-martini/martini"
	"net/http"
	"github.com/martini-contrib/render"
	"github.com/heckfer/fala-com-meu-carro/model"
	"github.com/bearchinc/datastore-model"
	"log"
)

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