package model_test

import (
	"testing"

	"github.com/renatoosaka/imersao-fullstack-fullcycle/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestModel_NewUser(t *testing.T) {
	name := "Jhon Doe"
	email := "jhondoe@email.com"
	invalidEmail := "invalid-email"

	user, err := model.NewUser(name, email)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(user.ID))
	require.Equal(t, user.Name, name)
	require.Equal(t, user.Email, email)

	_, noNameErr := model.NewUser("", email)
	require.NotNil(t, noNameErr)

	_, noEmailErr := model.NewUser(name, "")
	require.NotNil(t, noEmailErr)

	_, invalidEmailErr := model.NewUser(name, invalidEmail)
	require.NotNil(t, invalidEmailErr)
}
