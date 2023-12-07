package controller

import (
	"gin/config"
	"gin/model"
	"gin/shared/common"
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
		common.SendErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	task, err := t.taskUC.RegisterNewTask(payload)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	common.SendCreateResponse(ctx, task, "Created")
}

func (t *TaskController) listHandler(ctx *gin.Context) {
	tasks, err := t.taskUC.FindAllTask()
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	var response []interface{}
	for _, v := range tasks {
		response = append(response, v)
	}
	common.SendPagedResponse(ctx, response, "Ok")
}

func (t *TaskController) getByAuthorHandler(ctx *gin.Context) {
	author := ctx.Param("author")
	tasks, err := t.taskUC.FindTaskByAuthor(author)
	if err != nil {
		common.SendErrorResponse(ctx, http.StatusNotFound, "task with author ID "+author+" not found")
		return
	}
	common.SendSingleResponse(ctx, tasks, "Ok")
}

func (t *TaskController) Route() {
	t.rg.POST(config.TaskPost, t.createHandler)
	t.rg.GET(config.TaskGetList, t.listHandler)
	t.rg.GET(config.TaskGetByAuthor, t.getByAuthorHandler)
}

func NewTaskController(taskUC usecase.TaskUseCase, rg *gin.RouterGroup) *TaskController {
	return &TaskController{
		taskUC: taskUC,
		rg:     rg,
	}
}
