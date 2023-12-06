package usecase

import (
	"fmt"
	"gin/model"
	"gin/repository"
	"time"
)

type TaskUseCase interface {
	FindAllTask() ([]model.Task, error)
	RegisterNewTask(payload model.Task) (model.Task, error)
}
type taskUseCase struct {
	repo repository.TaskRepository
	//authorUC usecase.AuthorUseCase
}

func (t *taskUseCase) FindAllTask() ([]model.Task, error) {
	return t.repo.List()
}

func (t *taskUseCase) RegisterNewTask(payload model.Task) (model.Task, error) {
	//author, err := t.authorUC.FindAuthorById(payload.AuthorId)

	if payload.Title == "" || payload.Content == "" {
		return model.Task{}, fmt.Errorf("oppps, required fields")
	}
	payload.UpdatedAt = time.Now()
	task, err := t.repo.Create(payload)
	if err != nil {
		return model.Task{}, fmt.Errorf("oppps, failed to save data task :%v", err.Error())
	}
	return task, nil
}

func NewTaskUseCase(repo repository.TaskRepository) TaskUseCase {
	return &taskUseCase{repo: repo}
}
