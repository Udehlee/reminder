package routes

import (
	"github.com/Udehlee/payment-reminder/api/handler"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, h handler.Handler) {

	//home
	r.GET("/", h.Index)

	//user
	userRoute := r.Group("/user")
	{
		userRoute.POST("/register", h.Register)
		userRoute.POST("/login", h.Login)
		userRoute.GET("/", h.Index)

	}

	// TODO: route for reminders

}
