package server

import "github.com/gin-gonic/gin"

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

var GIN_ROUTER = gin.Default()

func InitHttpServer(port string) {
	GIN_ROUTER.Run(":" + port)
}
