package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
)

func (handler *Handler) GetReport(c *gin.Context) {
	var year int
	var month int
	year, err := strconv.Atoi(c.Request.URL.Query().Get("year"))
	if err != nil {
		panic(err)
		return
	}
	month, err = strconv.Atoi(c.Request.URL.Query().Get("month"))
	if err != nil {
		panic(err)
		return
	}
	link, err := handler.service.Report.GetReport(year, month)
	if err != nil {
		panic(err)
		return
	}
	if err := InitConfig(); err != nil {
		panic(err)
	}
	c.String(http.StatusOK, fmt.Sprintf("http://%s:%s/report?filename=%s", viper.Get("host"), viper.Get("port"), link))

}

func InitConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
