package handler
import (
	"github.com/go-martini/martini"
	"net/http"
	"github.com/martini-contrib/render"
	"github.com/heckfer/fala-com-meu-carro/model"
	"log"

	"github.com/drborges/appx"
)

func ListPostsHandler(c martini.Context, req *http.Request, r render.Render, appx *appx.Datastore) {

	var posts []*model.Post
	q := model.Posts.All()
	err := appx.Query(q).Results(&posts)

	if err != nil {
		log.Printf("Error: %+v", err)
		r.JSON(http.StatusInternalServerError, "Error")
	} else {
		r.JSON(http.StatusOK, posts)
	}
}

//var places []*models.Place
//err := appx.Query(models.Places.ByRecent()).
//StreamOf(models.Place{}).
//Each(models.Places.LoadBroadcasts(appx)).
//CollectAs(&places)
//
//if err != nil {
//logger.Infof("Error fetching places: Err: %+v", err)
//render.JSON(http.StatusInternalServerError, []*models.Place{})
//return
//}
//
//render.JSON(http.StatusOK, places)