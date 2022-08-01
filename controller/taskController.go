package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/sharelin-linpeng/easyRun/common/id"
	"github.com/sharelin-linpeng/easyRun/common/jsonutil"
	"github.com/sharelin-linpeng/easyRun/common/server"
	"github.com/sharelin-linpeng/easyRun/entity"
	"github.com/sharelin-linpeng/easyRun/repository"
)

func InitTaskRouter() {
	server.GIN_ROUTER.POST("/api/task", addTask)
	server.GIN_ROUTER.GET("/api/task/:id", findTaskById)
	server.GIN_ROUTER.GET("/api/task", findTaskList)
	server.GIN_ROUTER.DELETE("/api/task/:id", deleteTask)

}

func findTaskList(c *gin.Context) {
	app, err := repository.TaskService.QueryList()
	if err != nil {
		server.CreateError(c, err.Error())
		return
	}
	server.CreateSuccessData(c, "查询成功", app)
}

func addTask(c *gin.Context) {
	values, ok := c.GetPostFormArray("tasks")
	if !ok {
		log.Printf("addPubishInfo 参数异常\n")
		server.CreateError(c, "参数异常")
		return
	}

	tasks := make([]entity.Task, len(values))
	parentId := id.SNOW_FLAKE.NextVal()
	for i, v := range values {
		task := entity.Task{}
		jsonutil.Json2Obj(v, &task)
		task.Sequence = i
		if i == 0 {
			task.ParentId = 0
			task.Id = parentId
		} else {
			task.ParentId = parentId
			task.Id = id.SNOW_FLAKE.NextVal()

		}
		tasks[i] = task

	}

	for _, task := range tasks {
		if task.Type == entity.UPDATE_CODE {
			if task.CodeBranchId == "" {
				server.CreateError(c, "分支ID不能为空")
				return
			}
		} else if task.Type == entity.BUILD_CODE {
			if task.CodeBranchId == "" {
				server.CreateError(c, "分支ID不能为空")
				return
			}
			if task.BuildPath == "" {
				server.CreateError(c, "构建目录不能为空")
				return
			}
		} else if task.Type == entity.PUBLISH_APP {
			if task.CodeBranchId == "" {
				server.CreateError(c, "分支ID不能为空")
				return
			}
			if task.AppId == "" {
				server.CreateError(c, "应用信息不能为空")
				return
			}
		}
	}

	for _, task := range tasks {
		repository.TaskService.Add(task)
	}
	server.CreateSuccessData(c, "新增成功", tasks)

}

// 查找 byId
func findTaskById(c *gin.Context) {
	id := c.Param("id")

	app, err := repository.TaskService.QueryById(id)
	if err != nil {
		server.CreateError(c, err.Error())
		return
	}

	server.CreateSuccessData(c, "查询成功", app)

}

// 删除
func deleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := repository.TaskService.Delete(id); err != nil {
		server.CreateError(c, err.Error())
		return
	}
	server.CreateSuccess(c, "删除成功")

}
