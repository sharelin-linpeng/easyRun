package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sharelin-linpeng/easyRun/common/server"
	"github.com/sharelin-linpeng/easyRun/entity"
	"github.com/sharelin-linpeng/easyRun/repository"
)

func InitApplicationRouter() {
	server.GIN_ROUTER.POST("/application/add", addApplication)
	server.GIN_ROUTER.GET("/application/query/:id", findApplication)
	server.GIN_ROUTER.GET("/application/query", findApplicationList)
	server.GIN_ROUTER.POST("/application/update", updateApplication)
	server.GIN_ROUTER.DELETE("/application/delete/:id", deleteApplication)

}

// 新增应用
func addApplication(c *gin.Context) {
	app := entity.Application{}
	if err := c.ShouldBind(&app); err != nil {
		server.CreateError(c, "参数异常")
	}
	if err := repository.ApplicationService.Add(app); err != nil {
		server.CreateError(c, err.Error())
	} else {
		server.CreateSuccess(c, "添加应用成功")
	}

}

// 查找应用 byId
func findApplication(c *gin.Context) {
	id := c.Param("id")
	app, err := repository.ApplicationService.QueryById(id)
	if err != nil {
		server.CreateError(c, err.Error())
	} else {
		server.CreateSuccessData(c, "查询成功", app)
	}

}

// 查找应用列表
func findApplicationList(c *gin.Context) {
	app, err := repository.ApplicationService.QueryList()
	if err != nil {
		server.CreateError(c, err.Error())
	} else {
		server.CreateSuccessData(c, "查询成功", app)
	}

}

// 更新应用
func updateApplication(c *gin.Context) {
	app := entity.Application{}
	if err := c.ShouldBind(&app); err != nil {
		server.CreateError(c, "参数异常")
	}

	if err := repository.ApplicationService.Update(app); err != nil {
		server.CreateError(c, err.Error())
	} else {
		server.CreateSuccessData(c, "修改成功", app)
	}

}

// 删除 应用
func deleteApplication(c *gin.Context) {
	id := c.Param("id")
	if err := repository.ApplicationService.Delete(id); err != nil {
		server.CreateError(c, err.Error())
	} else {
		server.CreateSuccess(c, "删除成功")
	}

}
