package model
import "github.com/bearchinc/datastore-model"

type Car struct {
	db.Model `db:"Car"`
	Id    string `json:"id" db:"id"`
	Plate string `json: plate`
}

