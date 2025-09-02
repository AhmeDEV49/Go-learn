package main

import (
	"github.com/ahmedev49/go-clean-architecture/internal/handler"
	"github.com/ahmedev49/go-clean-architecture/internal/infrastructure"
	"github.com/ahmedev49/go-clean-architecture/internal/usecase"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	dsn := "postgres://usertest@localhost:5432/dbtest?sslmode=disable"
	db := infrastructure.NewPostgresDB(dsn)

	userRepo := infrastructure.NewUserRepositoryPostgres(db)

	userUc := usecase.NewUserUseCase(userRepo)

	userHandler := handler.NewUserHandler(userUc)

	r := gin.Default()

	r.POST("/users", userHandler.CreateUser)
	r.GET("/users/:id", userHandler.GetUserById)
	r.GET("/users", userHandler.GetAll)

	log.Println("Server running at http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
