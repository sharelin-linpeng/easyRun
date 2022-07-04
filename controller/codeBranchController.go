package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sharelin-linpeng/easyRun/common/server"
	"github.com/sharelin-linpeng/easyRun/entity"
	"github.com/sharelin-linpeng/easyRun/repository"
)

func InitCodeBranchRouter() {
	server.GIN_ROUTER.POST("/codeBranch", addCodeBranch)
	server.GIN_ROUTER.GET("/codeBranch/:id", findCodeBranchById)
	server.GIN_ROUTER.GET("/codeBranch", findCodeBranchList)
	server.GIN_ROUTER.PUT("/codeBranch/:id", updateCodeBranch)
	server.GIN_ROUTER.DELETE("/codeBranch/:id", deleteCodeBranch)

}

func addCodeBranch(c *gin.Context) {
	param := entity.CodeBranch{}
	if err := c.ShouldBind(&param); err != nil {
		log.Printf("addCodeBranch acc err %v\n", err)
		server.CreateError(c, "参数异常")
		return
	}
	if err := repository.CodeBranchService.Add(param); err != nil {
		server.CreateError(c, err.Error())
		return
	}
	server.CreateSuccess(c, "添加成功")

}

// 查找 byId
func findCodeBranchById(c *gin.Context) {
	id := c.Param("id")
	app, err := repository.CodeBranchService.QueryById(id)
	if err != nil {
		server.CreateError(c, err.Error())
		return
	}
	server.CreateSuccessData(c, "查询成功", app)

}

// 查找列表
func findCodeBranchList(c *gin.Context) {
	app, err := repository.CodeBranchService.QueryList()
	if err != nil {
		server.CreateError(c, err.Error())
		return
	}
	server.CreateSuccessData(c, "查询成功", app)

}

// 更新
func updateCodeBranch(c *gin.Context) {
	app := entity.CodeBranch{}
	if err := c.ShouldBind(&app); err != nil {
		log.Printf("updateCodeBranch acc err %v\n", err)
		server.CreateError(c, "参数异常")
		return
	}
	id := c.Param("id")
	app.Id = id

	if err := repository.CodeBranchService.Update(app); err != nil {
		server.CreateError(c, err.Error())
		return
	}
	server.CreateSuccessData(c, "修改成功", app)

}

// 删除
func deleteCodeBranch(c *gin.Context) {
	id := c.Param("id")
	if err := repository.CodeBranchService.Delete(id); err != nil {
		server.CreateError(c, err.Error())
		return
	}
	server.CreateSuccess(c, "删除成功")

}
