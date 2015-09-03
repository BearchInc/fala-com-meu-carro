package handler
import (
	"log"
	"github.com/heckfer/fala-com-meu-carro/model"
	"github.com/bearchinc/datastore-model"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

func ListPostsByCarPlateHandler(c martini.Context, req *http.Request, r render.Render, params martini.Params) {

	context := getAppengineContext(req)

	posts := model.Posts{}
	carPlate := params["plate"]
	err := db.NewDatastore(context).Query(db.From(new(model.Post)).Filter("CarPlate=", carPlate)).All(&posts)

	if err != nil {
		log.Printf("Error: %+v", err)
		r.JSON(500, "Error")
	} else {
		r.JSON(200, posts)
	}
}
