package controller

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sharelin-linpeng/easyRun/common/server"
	"github.com/sharelin-linpeng/easyRun/entity"
	"github.com/sharelin-linpeng/easyRun/repository"
)

func InitApplicationRouter() {
	server.GIN_ROUTER.POST("/api/application", addApplication)
	server.GIN_ROUTER.GET("/api/application/:id", findApplication)
	server.GIN_ROUTER.GET("/api/application", findApplicationList)
	server.GIN_ROUTER.PUT("/api/application/:id", updateApplication)
	server.GIN_ROUTER.DELETE("/api/application/:id", deleteApplication)

}

// 新增应用
func addApplication(c *gin.Context) {
	app := entity.Application{}
	if err := c.ShouldBind(&app); err != nil {
		log.Printf("addApplication acc err %v\n", err)
		server.CreateError(c, "参数异常")
		return
	}
	if err := repository.ApplicationService.Add(app); err != nil {
		server.CreateError(c, err.Error())
		return
	}
	server.CreateSuccess(c, "添加应用成功")

}

// 查找应用 byId
func findApplication(c *gin.Context) {
	id := c.Param("id")
	app, err := repository.ApplicationService.QueryById(id)
	if err != nil {
		server.CreateError(c, err.Error())
		return
	}
	server.CreateSuccessData(c, "查询成功", app)

}

// 查找应用列表
func findApplicationList(c *gin.Context) {
	app, err := repository.ApplicationService.QueryList()
	if err != nil {
		server.CreateError(c, err.Error())
		return
	}
	server.CreateSuccessData(c, "查询成功", app)

}

// 更新应用
func updateApplication(c *gin.Context) {
	id := c.Param("id")
	app := entity.Application{}
	if err := c.ShouldBind(&app); err != nil {
		log.Printf("updateApplication acc err %v\n", err)
		server.CreateError(c, "参数异常")
		return
	}
	app.Id, _ = strconv.Atoi(id)
	if err := repository.ApplicationService.Update(app); err != nil {
		server.CreateError(c, err.Error())
		return
	}
	server.CreateSuccessData(c, "修改成功", app)

}

// 删除 应用
func deleteApplication(c *gin.Context) {
	id := c.Param("id")
	if err := repository.ApplicationService.Delete(id); err != nil {
		server.CreateError(c, err.Error())
		return
	}
	server.CreateSuccess(c, "删除成功")

}
