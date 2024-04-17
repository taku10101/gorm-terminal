package routes

import (
	"gorm-terminal/internal/interfaces/handler"

	"github.com/gin-gonic/gin"
)

type Route struct {
	router      *gin.Engine
	userHandler *handler.UserHandler
}

// NewRoute はRouteのポインタを生成します。
func NewRoute(router *gin.Engine, userHandler *handler.UserHandler) *Route {
	return &Route{router, userHandler}
}


// Setup はルーティングを設定します。
func (r *Route) Setup() {
	router := gin.Default()

	router.GET("/users", r.userHandler.FindAll())
}