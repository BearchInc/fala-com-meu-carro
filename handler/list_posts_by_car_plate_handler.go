package handler
import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

func ListPostsByCarPlateHandler(c martini.Context, req *http.Request, r render.Render, params martini.Params) {

	//	context := getAppengineContext(req)
	//
	//	posts := []*model.Post{}
	//	carPlate := params["plate"]
	//	err := appx.NewDatastore(context).Query(db.From(new(model.Post)).Filter("CarPlate=", carPlate)).All(&posts)
	//
	//	if err != nil {
	//		log.Printf("Error: %+v", err)
	//		r.JSON(500, "Error")
	//	} else {
	//		r.JSON(200, posts)
	//	}
}
