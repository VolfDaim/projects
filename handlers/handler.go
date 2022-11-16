package handlers

import (
	"github.com/gin-gonic/gin"
	"projects/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (handler *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	user := router.Group("/user")
	{
		user.GET("/:id", handler.GetUser)
		user.POST("/create-user", handler.CreateUser)
		user.GET("users", handler.GetAllUsers)
		user.PUT(":id", handler.UpdateUser)
		user.DELETE(":id", handler.DeleteUser)
	}

	/*transaction := router.Group("/transaction")
	{
		transaction.GET(":id")
		transaction.PUT(":id")
	}
	*/
	return router
}
