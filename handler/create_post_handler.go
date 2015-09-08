package handler
import (
	"github.com/martini-contrib/render"
	"github.com/heckfer/fala-com-meu-carro/model"
	"strings"
	"github.com/drborges/appx"
	"log"
	"net/http"
"time"
)

func CreatePostHandler(r render.Render, post model.Post, appx *appx.Datastore) {

	response := model.Response{
		ErrorCode: http.StatusOK,
		Message: "",
		Data: nil,
	}

	post.CarPlate = strings.ToUpper(post.CarPlate)
	post.CreatedAt = time.Now()
	err := appx.Save(&post)

	if err != nil {
		log.Printf("Error: %+v", err)
		response.ErrorCode = http.StatusInternalServerError
		response.Message = "Some error happened"
	} else {
		response.Data = post
	}

	r.JSON(response.ErrorCode, response)
}
