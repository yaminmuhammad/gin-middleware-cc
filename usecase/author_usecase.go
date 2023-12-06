package usecase

import (
	"gin/model"
	"gin/repository"
)

type AuthorUseCase interface {
	FindAuthorByID(id string) (model.Author, error)
}

type authorUseCase struct {
	repo repository.AuthorRepository
}

func (a *authorUseCase) FindAuthorByID(id string) (model.Author, error) {
	return a.repo.Get(id)
}

func NewAuthorUseCase(repo repository.AuthorRepository) AuthorUseCase {
	return &authorUseCase{repo: repo}
}
