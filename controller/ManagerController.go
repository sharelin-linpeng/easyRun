package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sharelin-linpeng/easyRun/jsonutil"
	"github.com/sharelin-linpeng/easyRun/repository"
)

func InitHttpServer(port string) *gin.Engine {
	router := gin.Default()
	router.POST("/application/add", addApplication)
	router.Run(":" + port)
	return router
}

const (
	SUCCESS = 0
	ERROR   = 1
)

type ServerResponse struct {
	Status  int8        `json:"status"`
	Message string      `json:"message"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func CreateSuccess(c *gin.Context, message string) {
	receiver := ServerResponse{}
	receiver.Message = message
	receiver.Status = SUCCESS
	c.JSON(200, receiver)
}

func CreateSuccessData(c *gin.Context, message string, data interface{}) {
	receiver := ServerResponse{}
	receiver.Message = message
	receiver.Status = SUCCESS
	receiver.Data = data
	c.JSON(200, receiver)
}

func CreateError(c *gin.Context, message string) {
	receiver := ServerResponse{}
	receiver.Message = message
	receiver.Status = ERROR
	c.JSON(200, receiver)
}

func addApplication(c *gin.Context) {
	param := c.PostForm("param")
	app := repository.Application{}
	jsonutil.Json2Obj(param, &app)
	if err := repository.ApplicationService.Add(app); err != nil {
		CreateError(c, err.Error())
	} else {
		CreateSuccess(c, "添加应用成功")
	}

}
