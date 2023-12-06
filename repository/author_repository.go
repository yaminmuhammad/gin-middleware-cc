package repository

import (
	"database/sql"
	"fmt"
	"gin/config"
	"gin/model"
	"log"
)

type AuthorRepository interface {
	Get(id string) (model.Author, error)
}

type authorRepository struct {
	db *sql.DB
}

func (a *authorRepository) Get(id string) (model.Author, error) {
	fmt.Println("id:", id)
	var author model.Author
	err := a.db.QueryRow(config.SelectAuthorByID, id).Scan(
		&author.ID,
		&author.Name,
		&author.Email,
		&author.CreatedAt,
		&author.UpdatedAt,
	)
	if err != nil {
		log.Println("authorRepository.err:", err.Error())
		return model.Author{}, err
	}
	return author, nil
}

func NewAuthorRepository(db *sql.DB) AuthorRepository {
	return &authorRepository{db: db}
}
