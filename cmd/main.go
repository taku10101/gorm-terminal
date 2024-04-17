package main

import (
	"gorm-terminal/internal/application/service"
	"gorm-terminal/internal/domain/repository"
	"gorm-terminal/internal/interfaces/database"
	"gorm-terminal/internal/interfaces/handler"
	"gorm-terminal/internal/routes"

	"log"
)


func main() {
    db, err := database.InitializeDB("host=localhost user=postgres password=postgres dbname=postgres port=5412 sslmode=disable TimeZone=Asia/Tokyo")
    if err != nil {
        log.Fatal("Database initialization failed: ", err)
    }

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(*userRepository)
	userHandler := handler.NewUserHandler(*userService)

	routes :=routes.NewRoute(userHandler)
	routes.Setup()

	

}
