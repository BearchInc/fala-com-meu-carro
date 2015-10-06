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

func (user *User) IsValid() (bool, []string) {
	errors := []string{}

	if user.DisplayName == "" {
		errors = append(errors, "Nome de usuário não pode ser vazio.")
	}

	if user.Email == "" {
		errors = append(errors, "Email não pode ser vazio.")
	}

	if user.Password == "" {
		errors = append(errors, "Password não pode ser vazio.")
	}

	return len(errors) == 0, errors
}
