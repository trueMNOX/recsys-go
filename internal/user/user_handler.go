package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *UserServiec
}

func NewUserHanler(userService *UserServiec) *UserHandler{
	return &UserHandler{service: userService}
}

func(h *UserHandler) UserRoute(r *gin.RouterGroup){
	r.POST("/user/sign", h.SignNewUser)
}

func (h *UserHandler) SignNewUser(c *gin.Context){
	var input *User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The input is incorrect"})
		return
	}

	user , err := h.service.SignNewUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The input is incorrect"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}
