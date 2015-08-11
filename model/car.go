package model
import "github.com/bearchinc/datastore-model"

type Car struct {
	db.Model `db:"Car"`
	Plate string `json:"plate" form:"plate"`
}

