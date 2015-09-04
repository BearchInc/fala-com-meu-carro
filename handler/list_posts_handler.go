package handler
import (
	"github.com/go-martini/martini"
	"net/http"
	"github.com/martini-contrib/render"
	"github.com/heckfer/fala-com-meu-carro/model"
	"log"

	"github.com/drborges/appx"
)

func ListPostsHandler(c martini.Context, req *http.Request, r render.Render, appx *appx.Datastore) {

	var posts []*model.Post
	err := appx.Query(model.Posts.All()).Results(&posts)

	responseCode := http.StatusOK
	message := ""
	var data []*model.Post

	if err != nil {
		log.Printf("Error: %+v", err)

		responseCode = http.StatusInternalServerError
		message = "Some error happened"
		data = []*model.Post{}
	} else {
		data = posts
	}

	response := model.Response{
		ErrorCode:responseCode,
		Message: message,
		Data: data,
	}

	r.JSON(responseCode, response)
}
