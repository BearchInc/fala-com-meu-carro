package handler
import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"log"
	"github.com/heckfer/fala-com-meu-carro/model"
	"github.com/drborges/appx"
	"net/http"
	"appengine/datastore"
)

func ListPostsByCarPlateHandler(r render.Render, params martini.Params, appx *appx.Datastore) {

	carPlate := params["plate"]

	response := model.Response{
		ErrorCode: http.StatusOK,
		Message: []string{},
		Data: &[]*model.Post{},
	}

	err := appx.Query(model.Posts.AllByCarPlate(carPlate)).Results(response.Data)

	model.SetPostKeys(response.Data.([]*model.Post))

	if err != nil && err != datastore.Done {
		log.Printf("Error: %+v", err)
		response.ErrorCode = http.StatusInternalServerError
		response.Message = append(response.Message, err.Error())
	}

	r.JSON(200, response)
}
