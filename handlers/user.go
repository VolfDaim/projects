package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projects/models"
)

type GetAllLists struct {
	Data []models.User `json:"data"`
}

func (handler *Handler) GetUser(c *gin.Context) {

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

func (handler *Handler) UpdateUser(c *gin.Context) {

}

func (handler *Handler) DeleteUser(c *gin.Context) {

}
