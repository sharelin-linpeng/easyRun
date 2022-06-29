package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sharelin-linpeng/easyRun/common/server"
	"github.com/sharelin-linpeng/easyRun/entity"
	"github.com/sharelin-linpeng/easyRun/repository"
)

func InitPubishInfoRouter() {
	server.GIN_ROUTER.POST("/publishInfo/add", addPubishInfo)
	server.GIN_ROUTER.GET("/publishInfo/query/:id", findPubishInfoById)
	server.GIN_ROUTER.GET("/publishInfo/query", findPubishInfoList)
	server.GIN_ROUTER.POST("/publishInfo/update", updatePubishInfo)
	server.GIN_ROUTER.DELETE("/publishInfo/delete/:id", deletePubishInfo)

}

func addPubishInfo(c *gin.Context) {
	param := entity.PublishInfo{}
	if err := c.ShouldBind(&param); err != nil {
		server.CreateError(c, "参数异常")
	}
	if err := repository.PublishInfoService.Add(param); err != nil {
		server.CreateError(c, err.Error())
	} else {
		server.CreateSuccess(c, "添加成功")
	}

}

// 查找 byId
func findPubishInfoById(c *gin.Context) {
	id := c.Param("id")
	app, err := repository.PublishInfoService.QueryById(id)
	if err != nil {
		server.CreateError(c, err.Error())
	} else {
		server.CreateSuccessData(c, "查询成功", app)
	}

}

// 查找列表
func findPubishInfoList(c *gin.Context) {
	app, err := repository.PublishInfoService.QueryList()
	if err != nil {
		server.CreateError(c, err.Error())
	} else {
		server.CreateSuccessData(c, "查询成功", app)
	}

}

// 更新
func updatePubishInfo(c *gin.Context) {
	app := entity.PublishInfo{}
	if err := c.ShouldBind(&app); err != nil {
		server.CreateError(c, "参数异常")
	}

	if err := repository.PublishInfoService.Update(app); err != nil {
		server.CreateError(c, err.Error())
	} else {
		server.CreateSuccessData(c, "修改成功", app)
	}

}

// 删除
func deletePubishInfo(c *gin.Context) {
	id := c.Param("id")
	if err := repository.PublishInfoService.Delete(id); err != nil {
		server.CreateError(c, err.Error())
	} else {
		server.CreateSuccess(c, "删除成功")
	}

}