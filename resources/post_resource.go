package resources
import (
	"time"
	"github.com/heckfer/fala-com-meu-carro/model"
)

type PostResource struct {
	Id        string `json:"id"`
	CarPlate  string `json:"car_plate"`
	Message   string `json:"message"`
	UserName  string `json:"user_name"`
	CreatedAt time.Time `json:"created_at"`
}

func (this *PostResource) From(post model.Post) {
	this.CarPlate = post.CarPlate
	this.CreatedAt = post.CreatedAt
	this.Id = post.EncodedKey()
	this.Message = post.Message
	this.UserName = post.UserName
}

func FromPostResource(posts []*model.Post) []*PostResource {
	list := []*PostResource{}
	for _, p := range posts {
		postResource := PostResource{}
		postResource.From(*p)
		list = append(list, &postResource)
	}
	return list
}