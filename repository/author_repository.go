package repository

import (
	"database/sql"
	"gin/config"
	"gin/model"
)

type AuthorRepository interface {
	List() ([]model.Author, error)
	Create(payload model.Author) (model.Author, error)
}

type authorRepository struct {
	db *sql.DB
}

func (a *authorRepository) List() ([]model.Author, error) {
	var authors []model.Author
	rows, err := a.db.Query(config.SelectAuthorList)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var author model.Author
		err = rows.Scan(
			&author.ID,
			&author.Name,
			&author.Email,
			&author.Password,
			&author.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}
	return authors, nil
}

func (t *authorRepository) Create(payload model.Author) (model.Author, error) {
	var author model.Author
	err := t.db.QueryRow(config.InsertAuthor,
		payload.Name,
		payload.Email,
		payload.Password,
		payload.UpdatedAt).Scan(
		&author.ID,
		&author.CreatedAt,
	)
	if err != nil {
		return model.Author{}, err
	}
	author.Name = payload.Name
	author.Email = payload.Email
	author.Password = payload.Password
	author.UpdatedAt = payload.UpdatedAt
	return author, nil
}

func NewAuthorRepository(db *sql.DB) AuthorRepository {
	return &authorRepository{db: db}
}
