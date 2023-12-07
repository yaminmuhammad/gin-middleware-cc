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
	List() ([]model.Author, error)
}

type authorRepository struct {
	db *sql.DB
}

// List implements AuthorRepository.
func (a *authorRepository) List() ([]model.Author, error) {
	var authors []model.Author
	rows, err := a.db.Query(config.SelectAuthorList)
	if err != nil {
		log.Println("authorRepository.Query:", err.Error())
		return nil, err
	}
	for rows.Next() {
		var author model.Author
		var task model.Task
		err := rows.Scan(
			&author.ID,
			&author.Name,
			&author.Email,
			&author.CreatedAt,
			&author.UpdatedAt,
			&task.ID,
			&task.Title,
			&task.Content,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			log.Println("authorRepository.Rows.Next():", err.Error())
			return nil, err
		}
		// author.Tasks = append(author.Tasks, task)
		authors = append(authors, author)

	}
	return authors, nil

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
