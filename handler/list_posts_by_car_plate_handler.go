package handler
import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"log"
	"github.com/heckfer/fala-com-meu-carro/model"
	"github.com/drborges/appx"
	"net/http"
	"appengine/datastore"
	"github.com/heckfer/fala-com-meu-carro/resources"
	"github.com/heckfer/fala-com-meu-carro/middleware"
)

func ListPostsByCarPlateHandler(r render.Render, params martini.Params, appx *appx.Datastore, location middleware.RequestLocation) {
	carPlate := params["plate"]

	response := model.Response{
		ErrorCode: http.StatusOK,
		Message: []string{},
		Data: nil,
	}

	posts := []*model.Post{}
	err := appx.Query(model.Posts.AllByCarPlate(carPlate, location.Country)).Results(&posts)
	response.Data = resources.FromPostResource(posts)

	if err != nil && err != datastore.Done {
		log.Printf("Error: %+v", err)
		response.ErrorCode = http.StatusInternalServerError
		response.Message = append(response.Message, err.Error())
	}

	r.JSON(200, response)
}
