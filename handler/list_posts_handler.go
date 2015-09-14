package handler
import (
	"net/http"
	"github.com/martini-contrib/render"
	"github.com/heckfer/fala-com-meu-carro/model"
	"log"

	"github.com/drborges/appx"
	"appengine/datastore"
	"github.com/heckfer/fala-com-meu-carro/resources"
)

func ListPostsHandler(r render.Render, appx *appx.Datastore) {
	response := model.Response{
		ErrorCode: http.StatusOK,
		Message: []string{},
		Data: nil,
	}

	posts := []*model.Post{}
	err := appx.Query(model.Posts.All()).Results(&posts)

	response.Data = resources.FromPostResource(posts)

	if err != nil && err != datastore.Done {
		log.Printf("Error: %+v", err)
		response.ErrorCode = http.StatusInternalServerError
		response.Message = append(response.Message, err.Error())
	}

	r.JSON(200, response)
}
