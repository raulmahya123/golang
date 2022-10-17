package handler

import (
	"golang/helper"
	"golang/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse("avcooun gagal", http.StatusBadGateway, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("avcooun gagal", http.StatusBadGateway, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(newUser, "tokennn berhasil")
	response := helper.APIResponse("avcooun berhasil", http.StatusOK, "sukses", formatter)

	c.JSON(http.StatusOK, response)
}
