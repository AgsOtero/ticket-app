package main

import (
	"database/sql"
	"log"

	"github.com/AgsOtero/event-ticket-api/internal/adapters/db"
	"github.com/AgsOtero/event-ticket-api/internal/adapters/http"
	"github.com/AgsOtero/event-ticket-api/internal/core/services"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	dsn := "postgres://postgres:mysecretpassword@localhost:5432/tickets_db?sslmode=disable"

	dbConn, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("unable to conect to DB", err)
	}
	defer dbConn.Close()

	if err := dbConn.Ping(); err != nil {
		log.Fatal("unable to ping DB", err)
	}

	log.Println("Connected to DB")

	userPostgresRepo := db.NewPostgresUserRepository(dbConn)
	userService := services.NewUserService(userPostgresRepo)
	userHandler := http.NewUserHandler(userService)
	log.Println("User service initialized")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/users/register", userHandler.Register)
	r.GET("/users/:id", userHandler.GetById)
	r.Run()
}
