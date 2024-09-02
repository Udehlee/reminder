package handler

import (
	"net/http"

	models "github.com/Udehlee/payment-reminder/models/user"
	"github.com/Udehlee/payment-reminder/service"
	"github.com/Udehlee/payment-reminder/utils"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service service.Service
}

func NewHandler(svc service.Service) *Handler {
	return &Handler{
		service: svc,
	}
}

func (h *Handler) Index(c *gin.Context) {
	c.String(200, "Welcome Home")
}

// Register create new user info
func (h *Handler) Register(c *gin.Context) {
	var regReq models.CreateUserReq

	if err := c.ShouldBindJSON(&regReq); err != nil {
		utils.BadRequestErr(c, "Bad Request", "unable to read request", http.StatusBadRequest)
		return
	}

	user, err := h.service.CreateUser(regReq.FirstName, regReq.LastName, regReq.Email, regReq.Password)
	if err != nil {
		utils.BadRequestErr(c, "Bad Request", "Registration unsuccessful", http.StatusBadRequest)
		return
	}

	user.UserID = utils.GenerateUUID()

	createdUser := models.UserResponseInfo{
		UserID:    user.UserID,
		FullName:  user.FirstName + "" + user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Registration successful",
		"userinfo": createdUser,
	})

}

// Login logged in resgistered user
func (h *Handler) Login(c *gin.Context) {
	var loginReq models.LoginReq

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		utils.BadRequestErr(c, "Bad Request", "unable to read request", http.StatusBadRequest)
		return
	}

	user, err := h.service.CheckUser(loginReq.Email, loginReq.Password)
	if err != nil {
		utils.BadRequestErr(c, "Bad Request", "user not found", http.StatusBadRequest)
		return
	}

	logedinUser := models.UserResponseInfo{
		FullName:  user.FirstName + "" + user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "login successful",
		"userinfo": logedinUser,
	})
}
