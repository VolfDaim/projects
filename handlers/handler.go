package handlers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "projects/docs"
	"projects/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

// @title           Avito Tech Test
// @version         1.0
// @description		Микросервис для работы с балансом пользователей
// @host      		localhost:8081
// @BasePath  /
func (handler *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	user := router.Group("/user")
	{
		user.GET("/:id", handler.GetUser)
		user.GET("/:id/balance", handler.GetUserBalance)
		user.POST("/create-user", handler.CreateUser)
		user.GET("users", handler.GetAllUsers)
		user.DELETE("/:id", handler.DeleteUser)
	}

	transaction := router.Group("/transaction")
	{
		transaction.POST("/send", handler.Send)
		transaction.POST("/reserve", handler.Reserve)
		transaction.POST("/confirm", handler.Confirm)
	}

	report := router.Group("/report")
	{
		report.GET("/get-report", handler.GetReport)
		//report.GET("", handler.DownloadReport)
	}

	return router
}
