package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projects/models"
)

// Send godoc
// @Summary Send balance
// @Tags transaction
// @Description Зачислить средства на счет пользователя
// @ID send
// @Accept  json
// @Produce  json
// @Param input body models.User true "User data"
// @Router /transaction/send [post]
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

// Reserve godoc
// @Summary Reserve balance
// @Tags transaction
// @Description Заморозить средства пользователя
// @ID reserve
// @Accept  json
// @Produce  json
// @Param input body models.Transaction true "Transaction data"
// @Router /transaction/reserve [post]
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

// Confirm godoc
// @Summary Confirm transaction
// @Tags transaction
// @Description Подтвердить перевод и разморозить средства
// @ID confirm
// @Accept  json
// @Produce  json
// @Param input body models.Transaction true "Transaction data"
// @Router /transaction/confirm [post]
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
