package handler

import (
	"testing"

	models "github.com/Udehlee/reminder/models/user"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	m.Run()
}

type MockDB struct {
	mock.Mock
}

func (m *MockDB) SaveUser(user models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockDB) UserEmail(email string) (models.User, error) {
	args := m.Called(email)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockDB) UserPhoneNumber() ([]int, error) {
	args := m.Called()
	return nil, args.Error(0)

}
