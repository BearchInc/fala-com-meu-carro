package handler
import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"log"
	"github.com/heckfer/fala-com-meu-carro/model"
	"github.com/drborges/appx"
	"net/http"
)

func ListPostsByCarPlateHandler(r render.Render, params martini.Params, appx *appx.Datastore) {

	carPlate := params["plate"]

	var posts []*model.Post
	err := appx.Query(model.Posts.AllByCarPlate(carPlate)).Results(&posts)

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
