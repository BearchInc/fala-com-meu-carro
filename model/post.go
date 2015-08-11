package model
import "github.com/bearchinc/datastore-model"

type Post struct {
	db.Model `db:"Post"`
	Id       string `json:"id" db:"id"`
	CarPlate string `json: car_plate`
	Message  string `json: message`
}