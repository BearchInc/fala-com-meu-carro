package handler
import (
	"net/http"
	"github.com/martini-contrib/render"
	"github.com/heckfer/fala-com-meu-carro/model"
	"log"

	"github.com/drborges/appx"
)

func ListPostsHandler(r render.Render, appx *appx.Datastore) {

	var posts []*model.Post
	err := appx.Query(model.Posts.All()).Results(&posts)

	response := model.Response{
		ErrorCode: http.StatusOK,
		Message: "",
		Data: []*model.Post{},
	}

	if err != nil {
		log.Printf("Error: %+v", err)
		response.ErrorCode = http.StatusInternalServerError
		response.Message = "Some error happened"
	} else {
		response.Data = posts
	}

	r.JSON(response.ErrorCode, response)
}
