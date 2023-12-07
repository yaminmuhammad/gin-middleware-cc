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
	FindTaskByAuthorID(authorID string) ([]model.Task, error)
}

type taskUseCase struct {
	repo     repository.TaskRepository
	authorUC AuthorUseCase
}

// FindTaskByAuthorID implements TaskUseCase.
func (t *taskUseCase) FindTaskByAuthorID(authorID string) ([]model.Task, error) {
	tasks, err := t.repo.GetByAuthor(authorID)
	if err != nil {
		return nil, fmt.Errorf("failed to find tasks for author with ID %s: %v", authorID, err)
	}
	return tasks, nil
}

func (t *taskUseCase) FindAllTask() ([]model.Task, error) {
	return t.repo.List()
}

func (t *taskUseCase) RegisterNewTask(payload model.Task) (model.Task, error) {
	fmt.Println("payload.AuthorID:", payload.AuthorId)
	author, err := t.authorUC.FindAuthorByID(payload.AuthorId)
	fmt.Println("author:", author)
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
