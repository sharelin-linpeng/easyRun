package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sharelin-linpeng/easyRun/common/server"
	"github.com/sharelin-linpeng/easyRun/entity"
	"github.com/sharelin-linpeng/easyRun/repository"
)

func InitMachineRouter() {
	server.GIN_ROUTER.POST("/machine", addMachine)
	server.GIN_ROUTER.GET("/machine/:id", findMachineById)
	server.GIN_ROUTER.GET("/machine", findMachineList)
	server.GIN_ROUTER.PUT("/machine/:id", updateMachine)
	server.GIN_ROUTER.DELETE("/machine/:id", deleteMachine)

}

func addMachine(c *gin.Context) {
	param := entity.Machine{}
	if err := c.ShouldBind(&param); err != nil {
		log.Printf("addMachine acc err %v\n", err)
		server.CreateError(c, "参数异常")
		return
	}
	if err := repository.MachineService.Add(param); err != nil {
		server.CreateError(c, err.Error())
		return
	}
	server.CreateSuccess(c, "添加成功")

}

// 查找 byId
func findMachineById(c *gin.Context) {
	id := c.Param("id")
	app, err := repository.MachineService.QueryById(id)
	if err != nil {
		server.CreateError(c, err.Error())
		return
	}
	server.CreateSuccessData(c, "查询成功", app)

}

// 查找列表
func findMachineList(c *gin.Context) {
	app, err := repository.MachineService.QueryList()
	if err != nil {
		server.CreateError(c, err.Error())
		return
	}
	server.CreateSuccessData(c, "查询成功", app)

}

// 更新
func updateMachine(c *gin.Context) {
	app := entity.Machine{}
	if err := c.ShouldBind(&app); err != nil {
		log.Printf("updateMachine acc err %v\n", err)
		server.CreateError(c, "参数异常")
	}
	id := c.Param("id")
	app.Id = id

	if err := repository.MachineService.Update(app); err != nil {
		server.CreateError(c, err.Error())
		return
	}
	server.CreateSuccessData(c, "修改成功", app)

}

// 删除
func deleteMachine(c *gin.Context) {
	id := c.Param("id")
	if err := repository.MachineService.Delete(id); err != nil {
		server.CreateError(c, err.Error())
		return
	}
	server.CreateSuccess(c, "删除成功")

}
