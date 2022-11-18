package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projects/models"
	"strconv"
)

type GetAllLists struct {
	Data []models.User `json:"data"`
}

func (handler *Handler) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid type of user id")
		return
	}

	user, err := handler.service.User.GetUser(id)
	if err != nil {
		c.JSON(http.StatusOK, "User is not exist")
		return
	}
	c.JSON(http.StatusOK, user)
}

func (Handler *Handler) GetUserBalance(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, "Invalid type of user id")
	}

	balance, err := Handler.service.User.GetUserBalance(id)
	if err != nil {
		c.JSON(http.StatusOK, "User is not exist")
		return
	}
	c.JSON(http.StatusOK, balance)
}

func (handler *Handler) CreateUser(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		panic(err)
		return
	}
	id, err := handler.service.User.Create(input)
	if err != nil {
		panic(err)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (handler *Handler) GetAllUsers(c *gin.Context) {

	lists, err := handler.service.User.GetAll()
	if err != nil {
		panic(err)
		return
	}
	c.JSON(http.StatusOK, GetAllLists{Data: lists})
}

func (handler *Handler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusOK, "Invalid type of user id")
		return
	}

	err = handler.service.User.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusOK, "User is not exist")
		return
	}
	c.JSON(http.StatusOK, `json:"ok"`)
}
