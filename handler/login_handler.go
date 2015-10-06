package handler

import (
	"log"

	"net/http"

	"github.com/drborges/appx"
	"github.com/heckfer/fala-com-meu-carro/model"
	"github.com/martini-contrib/render"
)

func LoginHandler(r render.Render, user model.User, appx *appx.Datastore) {
	response := model.Response{
		ErrorCode: http.StatusOK,
		Message:   []string{},
		Data:      nil,
	}

	isValid, validationErr := user.IsValid()

	if !isValid {
		response.ErrorCode = http.StatusBadRequest
		response.Message = validationErr
	} else {
		err := appx.Save(&user)

		if err != nil {
			log.Printf("Error: %v", err)
			response.ErrorCode = http.StatusInternalServerError
			response.Message = append(response.Message, err.Error())
		} else {
			response.Data = user
		}
	}

	r.JSON(200, response)

}
