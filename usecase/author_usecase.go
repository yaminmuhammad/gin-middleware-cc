package usecase

import (
	"fmt"
	"gin/model"
	"gin/repository"
	"time"
)

type AuthorUseCase interface {
	FindAllAuthor() ([]model.Author, error)
	RegisterNewAuthor(payload model.Author) (model.Author, error)
}

type authorUseCase struct {
	repo repository.AuthorRepository
	//authorUC usecase.AuthorUseCase
}

func (t *authorUseCase) FindAllAuthor() ([]model.Author, error) {
	return t.repo.List()
}

func (t *authorUseCase) RegisterNewAuthor(payload model.Author) (model.Author, error) {
	//author, err := t.authorUC.FindAuthorById(payload.AuthorId)

	if payload.Name == "" || payload.Email == "" {
		return model.Author{}, fmt.Errorf("oppps, required fields")
	}
	payload.UpdatedAt = time.Now()
	author, err := t.repo.Create(payload)
	if err != nil {
		return model.Author{}, fmt.Errorf("oppps, failed to save data author :%v", err.Error())
	}
	return author, nil
}

func NewAuthorUseCase(repo repository.AuthorRepository) AuthorUseCase {
	return &authorUseCase{repo: repo}
}
