package service

import (
	"be15/clean/features/users"
	"be15/clean/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	repo := new(mocks.UserData)
	returnData := []users.Core{
		{Id: 1, Name: "budi", Email: "budi@mail.com", Password: "qwerty", Address: "Surabaya", Role: "user"},
	}

	t.Run("Success Get All Users", func(t *testing.T) {
		repo.On("SelectAll").Return(returnData, nil).Once()

		srv := New(repo)
		response, err := srv.GetAll()
		assert.Nil(t, err)
		assert.Equal(t, returnData[0].Name, response[0].Name)
		repo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	repo := new(mocks.UserData)
	// returnData := []users.Core{
	// 	{Id: 1, Name: "budi", Email: "budi@mail.com", Password: "qwerty", Address: "Surabaya", Role: "user"},
	// }

	t.Run("Success Insert Users", func(t *testing.T) {
		inputData := users.Core{Id: 1, Name: "budi", Email: "budi@mail.com", Password: "qwerty", Address: "Surabaya", Role: "user"}
		repo.On("Insert", inputData).Return(nil).Once()

		srv := New(repo)
		err := srv.Create(inputData)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})

	t.Run("Failed when func Insert return error", func(t *testing.T) {
		inputData := users.Core{Id: 1, Name: "budi", Email: "budi@mail.com", Password: "qwerty", Address: "Surabaya", Role: "user"}
		repo.On("Insert", inputData).Return(errors.New("error insert data")).Once()

		srv := New(repo)
		err := srv.Create(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, "error insert data", err.Error())
		repo.AssertExpectations(t)
	})

	t.Run("Failed validate", func(t *testing.T) {
		inputData := users.Core{Id: 1, Name: "budi", Password: "qwerty", Address: "Surabaya", Role: "user"}
		// repo.On("Insert", inputData).Return(errors.New("error insert data")).Once()

		srv := New(repo)
		err := srv.Create(inputData)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}
