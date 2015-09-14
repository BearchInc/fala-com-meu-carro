package handler
import (
	"github.com/martini-contrib/render"
	"github.com/heckfer/fala-com-meu-carro/middleware"
	"net/http"
)

type AppConfig struct {
	Location string `json:"location"`
}

func LoginHandler(r render.Render, location middleware.RequestLocation) {
	config := AppConfig{
		Location: location.Country,
	}
	r.JSON(http.StatusOK, config)
}