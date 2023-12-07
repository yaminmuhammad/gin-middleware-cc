package usecase

import (
	"gin/model"
	"gin/repository"
)

type AuthorUseCase interface {
	FindAllAuthor() ([]model.Author, error)
	FindAuthorByID(id string) (model.Author, error)
}

type authorUseCase struct {
	repo repository.AuthorRepository
}

func (a *authorUseCase) FindAllAuthor() ([]model.Author, error) {
	return a.repo.List()
}

func (a *authorUseCase) FindAuthorByID(id string) (model.Author, error) {
	return a.repo.Get(id)
}

func NewAuthorUseCase(repo repository.AuthorRepository) AuthorUseCase {
	return &authorUseCase{repo: repo}
}
