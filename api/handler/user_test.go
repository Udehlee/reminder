package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Udehlee/reminder/internals"
	models "github.com/Udehlee/reminder/models/user"
	"github.com/Udehlee/reminder/service"
	"github.com/Udehlee/reminder/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	reqUser := models.CreateUserReq{
		FirstName:   "Peter",
		LastName:    "Ada",
		Email:       "ada.example@email.com",
		Password:    "password",
		PhoneNumber: "+12345627",
	}

	// // Set up the MockDB
	mockDB := new(MockDB)
	mockDB.On("SaveUser", mock.MatchedBy(func(user models.User) bool {
		return user.FirstName == reqUser.FirstName &&
			user.LastName == reqUser.LastName &&
			user.Email == reqUser.Email
	})).Return(nil)

	//initialize dependencies
	logger := internals.NewLogger()
	scheduler := internals.NewScheduler()

	// Create a real Service, but with the mocked DB
	svc := *service.NewService(mockDB, logger, scheduler)
	h := NewHandler(svc)

	r := gin.Default()
	r.POST("/register", h.Register)

	userJson, err := json.Marshal(reqUser)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(userJson))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rc := httptest.NewRecorder()
	r.ServeHTTP(rc, req)

	assert.Equal(t, http.StatusOK, rc.Code)
	fmt.Println("Response Body:", rc.Body.String()) // for debugging purposes

	resp := gin.H{}
	err = json.Unmarshal(rc.Body.Bytes(), &resp)
	assert.NoError(t, err)

	assert.Equal(t, "Registration successful", resp["message"])
	assert.NotEmpty(t, resp["userinfo"])

	mockDB.AssertExpectations(t)

}

func TestLogin(t *testing.T) {
	loginReq := models.LoginReq{
		Email:    "ada.example@email.com",
		Password: "password",
	}

	// Hash the known password
	hashedPassword, err := utils.HashPassword(loginReq.Password)
	assert.NoError(t, err)

	expectedUser := models.User{
		UserID:   uuid.New().String(),
		Email:    loginReq.Email,
		Password: hashedPassword, // Store the hashed password
	}

	mockDB := new(MockDB)
	mockDB.On("UserEmail", mock.MatchedBy(func(email string) bool {
		return email == expectedUser.Email
	})).Return(expectedUser, nil)

	//initialize dependencies
	logger := internals.NewLogger()
	scheduler := internals.NewScheduler()

	// Create a real Service, but with the mocked DB
	svc := *service.NewService(mockDB, logger, scheduler)
	h := NewHandler(svc)

	r := gin.Default()
	r.POST("/login", h.Login)

	userJson, err := json.Marshal(loginReq)
	assert.NoError(t, err)

	req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(userJson))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	rc := httptest.NewRecorder()
	r.ServeHTTP(rc, req)

	assert.Equal(t, http.StatusOK, rc.Code)
	fmt.Println("Response Body:", rc.Body.String()) // for debugging purposes

	resp := gin.H{}
	err = json.Unmarshal(rc.Body.Bytes(), &resp)
	assert.NoError(t, err)

	assert.Equal(t, "login successful", resp["message"])
	assert.NotEmpty(t, resp["userinfo"])

	mockDB.AssertExpectations(t)
}
