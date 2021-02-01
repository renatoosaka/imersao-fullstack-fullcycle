package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

//User Define a estrutura do model
type User struct {
	Base  `valid:"-"`
	Name  string `json:"name" valid:"notnull"`
	Email string `json:"email" valid:"notnull"`
}

func (user *User) isValid() error {
	if _, err := govalidator.ValidateStruct(user); err != nil {
		return err
	}

	return nil
}

//NewUser Realiza a criação de um novo usuário ou retorna um erro
func NewUser(name, email string) (*User, error) {
	user := User{
		Name:  name,
		Email: email,
	}

	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err := user.isValid(); err != nil {
		return nil, err
	}

	return &user, nil
}
