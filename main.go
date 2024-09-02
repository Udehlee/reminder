package main

import (
	"log"

	"github.com/Udehlee/payment-reminder/api/handler"
	"github.com/Udehlee/payment-reminder/api/routes"
	"github.com/Udehlee/payment-reminder/db/db"
	"github.com/Udehlee/payment-reminder/internals"
	"github.com/Udehlee/payment-reminder/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	scheduler := internals.NewScheduler()
	logger := internals.NewLogger()

	config := db.InitLoadConfig()
	conn, err := db.InitPG(config)
	if err != nil {
		log.Fatal(err)
	}

	pgdb := db.NewPgDB(conn)
	svc := service.NewService(pgdb, logger, scheduler)
	svc.ScheduleTasks()

	h := handler.NewHandler(*svc)
	routes.SetupRoutes(r, *h)

	if err := r.Run(); err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
