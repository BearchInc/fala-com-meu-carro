package handler
import (
	"github.com/go-martini/martini"
	"net/http"
	"github.com/martini-contrib/render"
	"github.com/heckfer/fala-com-meu-carro/model"
	"strings"
	"github.com/drborges/appx"
)

func CreatePostHandler(c martini.Context, req *http.Request, r render.Render, post model.Post, appx *appx.Datastore) {

	post.CarPlate = strings.ToUpper(post.CarPlate)

	//	err := db.NewDatastore(context).Create(&post)
	//	if err != nil {
	//		log.Printf("Error: %+v", err)
	//		r.JSON(500, "Error")
	//	} else {
	//		r.JSON(200, post)
	//	}
}
