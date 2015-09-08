package model
import (
	"github.com/drborges/appx"
	"appengine/datastore"
	"time"
)

type Post struct {
	appx.Model

	CarPlate  string `json:"car_plate" form:"car_plate"`
	Message   string `json:"message" form:"message"`
	UserId    string `json:"user_id" form:"user_id"`
	UserName  string `json:"user_name" form:"user_name"`
	CreatedAt time.Time `json:"created_at"`
}

func (post *Post) KeySpec() *appx.KeySpec {
	return &appx.KeySpec{
		Kind:       "Post",
		Incomplete: true,
	}
}

var Posts = struct {
	All           func() *datastore.Query
	AllByCarPlate func(string) *datastore.Query
}{
	All: func() *datastore.Query {
		return datastore.NewQuery(new(Post).KeySpec().Kind).Order("-CreatedAt")
	},
	AllByCarPlate: func(carPlate string) *datastore.Query {
		return datastore.NewQuery(new(Post).KeySpec().Kind).Filter("CarPlate=", carPlate)
	},
}
