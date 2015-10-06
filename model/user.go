package model

import "github.com/drborges/appx"

type User struct {
	appx.Model
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

func (user *User) KeySpec() *appx.KeySpec {
	return &appx.KeySpec{
		Kind:       "User",
		Incomplete: true,
	}
}
