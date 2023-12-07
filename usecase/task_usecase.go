package usecase

import (
	"fmt"
	"gin/model"
	"gin/repository"
	"gin/shared/shared_model"
	"time"
)

type TaskUseCase interface {
	FindAllTask(page, size int) ([]model.Task, shared_model.Paging, error)
	FindTaskByAuthor(author string) ([]model.Task, error)
	RegisterNewTask(payload model.Task) (model.Task, error)
}

type taskUseCase struct {
	repo     repository.TaskRepository
	authorUC AuthorUseCase
}

func (t *taskUseCase) FindTaskByAuthor(author string) ([]model.Task, error) {
	return t.repo.GetByAuthor(author)
}

func (t *taskUseCase) FindAllTask(page, size int) ([]model.Task, shared_model.Paging, error) {
	return t.repo.List(page, size)
}

func (t *taskUseCase) RegisterNewTask(payload model.Task) (model.Task, error) {
	_, err := t.authorUC.FindAuthorByID(payload.AuthorId)
	if err != nil {
		return model.Task{}, fmt.Errorf("author with ID %s not found", payload.AuthorId)
	}
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

func NewTaskUseCase(repo repository.TaskRepository, authorUC AuthorUseCase) TaskUseCase {
	return &taskUseCase{repo: repo, authorUC: authorUC}
}
