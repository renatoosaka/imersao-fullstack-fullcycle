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

	user, err := model.NewUser(name, email)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(user.ID))
	require.Equal(t, user.Name, name)
	require.Equal(t, user.Email, email)

	_, errn := model.NewUser("", "")
	require.NotNil(t, errn)

}
