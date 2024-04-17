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

func (h *UserHandler) FindAll(c *gin.Context) ([]model.User, error) {
	users, err := h.userService.FindAll()

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (h *UserHandler) FindByID(c *gin.Context)(*model.User, error) {
	user_id := c.Param("id")

	int_id, err := strconv.Atoi(user_id)
	if err != nil {
		return nil, err
	}
	user, err := h.userService.FindByID(uint(int_id))
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (h *UserHandler) Create(c *gin.Context) error {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		return err
	}
	if err := h.userService.Create(&user); err != nil {
		return err
	}
	c.JSON(http.StatusOK, gin.H{"user": user})


	return nil

}

func (h *UserHandler) Update(c *gin.Context) error {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		return err
	}
	if err := h.userService.Update(&user); err != nil {
		
		return err
	}
	return nil
}
func (h *UserHandler) Delete(c *gin.Context) error {
	user_id := c.Param("id")
	int_id, err := strconv.Atoi(user_id)
	if err != nil {
		return err
	}
	if err := h.userService.Delete(uint(int_id)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})	

	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
	return nil
}