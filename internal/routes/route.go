package routes

import (
	"gorm-terminal/internal/interfaces/handler"

	"github.com/gin-gonic/gin"
)

type Route struct {
	userHandler *handler.UserHandler
}



// NewRoute はRouteのポインタを生成します。
func NewRoute( userHandler *handler.UserHandler) *Route {
	return &Route{userHandler}
}


// Setup はルーティングを設定します。
func (r *Route) Setup() {
	router := gin.Default()

	router.GET("/users", r.userHandler.FindAll)
	router.GET("/users/:id", r.userHandler.FindByID)
	router.POST("/users", r.userHandler.Create)
	router.PUT("/users/:id", r.userHandler.Update)
	router.DELETE("/users/:id", r.userHandler.Delete)

	router.Run(":8080")
}