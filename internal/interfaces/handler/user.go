package handler

import (
	"gorm-terminal/internal/application/service"
	"net/http"
	"strconv"

	"gorm-terminal/internal/domain/model"

	"github.com/gin-gonic/gin"
)


type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) FindAll(ctx *gin.Context) {
	users, err := h.userService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})	
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"users": users})
}

func (h *UserHandler) FindByID(ctx *gin.Context) {
	user_id := ctx.Param("id")
	int_id, err := strconv.Atoi(user_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.userService.FindByID(uint(int_id))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})	
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})	
}

func (h *UserHandler) Create(c *gin.Context)  {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.userService.Create(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})	
		return
	}
	c.JSON(http.StatusOK, gin.H{"created": user})
}

func (h *UserHandler) Update(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.userService.Update(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})	
		return
	}
	c.JSON(http.StatusOK, gin.H{"updated": user})
}

func (h *UserHandler) Delete(c *gin.Context) {
	user_id := c.Param("id")
	int_id, err := strconv.Atoi(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.userService.Delete(uint(int_id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})	
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": user_id})
}