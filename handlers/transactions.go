package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projects/models"
)

func (handler *Handler) Send(c *gin.Context) {
	var input models.User
	if err := c.BindJSON(&input); err != nil {
		panic(err)
		return
	}
	user, err := handler.service.Transaction.Send(input.ID, input.Balance)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":     user.ID,
		"amount": user.Balance,
	})
}

func (handler *Handler) Reserve(c *gin.Context) {
	var input models.Transaction
	if err := c.BindJSON(&input); err != nil {
		panic(err)
		return
	}
	err := handler.service.Transaction.Reserve(input.ID, input.IDService, input.IDOrder, input.Balance)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":         input.ID,
		"id_service": input.IDService,
		"id_order":   input.IDOrder,
		"balance":    input.Balance,
	})

}

// GetUser godoc
// @Summary Get user
// @Tags user
// @Description create user
// @Router /user [get]
func (handler *Handler) Confirm(c *gin.Context) {
	var input models.Transaction
	if err := c.BindJSON(&input); err != nil {
		panic(err)
		return
	}
	err := handler.service.Transaction.Confirm(input.ID, input.IDService, input.IDOrder, input.Balance)
	if err != nil {
		panic(err)
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id":         input.ID,
		"id_service": input.IDService,
		"id_order":   input.IDOrder,
		"balance":    input.Balance,
	})
}
