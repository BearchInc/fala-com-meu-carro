package model
import (
	"github.com/drborges/appx"
	"appengine/datastore"
)

type Post struct {
	appx.Model
	CarPlate string `json:"car_plate" form:"car_plate"`
	Message  string `json:"message" form:"message"`
}

func (post *Post) KeySpec() *appx.KeySpec {
	return &appx.KeySpec{
		Kind:       "Post",
		Incomplete: true,
	}
}

var Posts = struct {
	All func() *datastore.Query
}{
	All: func() *datastore.Query {
		return datastore.NewQuery(new(Post).KeySpec().Kind)
	},
}
