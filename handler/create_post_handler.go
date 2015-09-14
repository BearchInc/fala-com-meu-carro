package handler
import (
	"github.com/martini-contrib/render"
	"github.com/heckfer/fala-com-meu-carro/model"
	"strings"
	"github.com/drborges/appx"
	"log"
	"net/http"
	"time"
	"github.com/heckfer/fala-com-meu-carro/resources"
	"github.com/heckfer/fala-com-meu-carro/middleware"
)

func CreatePostHandler(r render.Render, post model.Post, appx *appx.Datastore, location middleware.RequestLocation) {

	response := model.Response{
		ErrorCode: http.StatusOK,
		Message: []string{},
		Data: nil,
	}

	post.CarPlate = strings.ToUpper(post.CarPlate)
	post.Country = location.Country
	post.CreatedAt = time.Now()

	isValid, validationErr := post.IsValid(location.Country)

	if !isValid {
		response.ErrorCode = http.StatusBadRequest
		response.Message = validationErr
	} else {
		err := appx.Save(&post)

		if err != nil {
			log.Printf("Error: %+v", err)
			response.ErrorCode = http.StatusInternalServerError
			response.Message = append(response.Message, err.Error())
		} else {
			postResource := &resources.PostResource{}
			postResource.From(post)
			response.Data = postResource
		}
	}

	r.JSON(200, response)
}

