package model

import "github.com/drborges/appx"

type User struct {
	appx.Model
	DisplayName string
	Email       string
	Password    string
}
