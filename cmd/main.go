package main

import (
	"gorm-terminal/internal/application/service"
	"gorm-terminal/internal/domain/repository"
	"gorm-terminal/internal/interfaces/database"
	"gorm-terminal/internal/interfaces/handler"
	"log"
)


func main() {
    db, err := database.InitializeDB("your_connection_string")
    if err != nil {
        log.Fatal("Database initialization failed: ", err)
    }

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(*userRepository)
	userHandler := handler.NewUserHandler(*userService)

}
