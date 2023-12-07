package controller

import (
	"gin/model"
	"gin/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskUC usecase.TaskUseCase
	rg     *gin.RouterGroup
}

func (t *TaskController) createHandler(ctx *gin.Context) {
	var payload model.Task
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	task, err := t.taskUC.RegisterNewTask(payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Ok",
		"data":    task,
	})
}

func (t *TaskController) listHandler(ctx *gin.Context) {
	tasks, err := t.taskUC.FindAllTask()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Ok",
		"data":    tasks,
	})
}

func (t *TaskController) Route() {
	t.rg.POST("/tasks", t.createHandler)
	t.rg.GET("/tasks", t.listHandler)
}

func NewTaskController(taskUC usecase.TaskUseCase, rg *gin.RouterGroup) *TaskController {
	return &TaskController{
		taskUC: taskUC,
		rg:     rg,
	}
}
