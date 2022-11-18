package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	_ "projects/docs"
	"projects/models"
	"strconv"
)

type GetAllLists struct {
	Data []models.User `json:"data"`
}

// GetUser godoc
// @Summary Get User
// @Tags user
// @Description Get User
// @ID get-user
// @Produce  json
// @Param id path int true "user id"
// @Router /user/{id} [get]
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

// GetUserBalance godoc
// @Summary Get User Balance
// @Tags user_balance
// @Description Get User Balance
// @ID get-user-balance
// @Accept  json
// @Produce  json
// @Param id path int true "user id"
// @Router /user/{id}/balance [get]
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

// CreateUser godoc
// @Summary Create User
// @Tags user
// @Description Create User
// @ID create-user
// @Accept  json
// @Produce  json
// @Param input body models.User true "User data"
// @Router /user/create-user [post]
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

// GetAllUsers godoc
// @Summary Get All Users
// @Tags user
// @Description Get All Users
// @ID get-all-user
// @Produce  json
// @Router /user/users [get]
func (handler *Handler) GetAllUsers(c *gin.Context) {

	lists, err := handler.service.User.GetAll()
	if err != nil {
		panic(err)
		return
	}
	c.JSON(http.StatusOK, GetAllLists{Data: lists})
}

// DeleteUser godoc
// @Summary Delete User
// @Tags user
// @Description Delete User
// @ID Delete-user
// @Accept  json
// @Produce  json
// @Param id path int true "user id"
// @Router /user/{id} [delete]
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
