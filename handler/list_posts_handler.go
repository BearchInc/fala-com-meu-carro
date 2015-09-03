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
	q := model.Posts.All()
	err := appx.Query(q).Results(&posts)

	if err != nil {
		log.Printf("Error: %+v", err)

		response := model.Response{
			ErrorCode:http.StatusInternalServerError,
			Message: "Some error happened",
			Data: []*model.Post{},
		}

		r.JSON(http.StatusInternalServerError, response)
	} else {

		response := model.Response{
			ErrorCode:http.StatusOK,
			Message: "",
			Data: posts,
		}

		r.JSON(http.StatusOK, response)
	}
}
