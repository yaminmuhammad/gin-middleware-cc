package repository

import (
	"database/sql"
	"gin/config"
	"gin/model"
	"log"
)

type AuthorRepository interface {
	List() ([]model.Author, error)
	Get(id string) (model.Author, error)
}

type authorRepository struct {
	db *sql.DB
}

func (a *authorRepository) List() ([]model.Author, error) {
	var authors []model.Author
	rows, err := a.db.Query(config.SelectAuthorList)
	if err != nil {
		log.Println("authorRepository.List.Query:", err.Error())
		return nil, err
	}
	for rows.Next() {
		var author model.Author
		err := rows.Scan(&author.ID, &author.Name, &author.Email, &author.CreatedAt, &author.UpdatedAt)
		if err != nil {
			log.Println("authorRepository.List.rows.Next:", err.Error())
			return nil, err
		}
		tasksRows, err := a.db.Query(config.SelectAuthorWithTask, author.ID)
		if err != nil {
			log.Println("authorRepository.Get.Query:", err.Error())
			return nil, err
		}
		for tasksRows.Next() {
			var task model.Task
			err := tasksRows.Scan(&task.ID, &task.Title, &task.Content, &task.CreatedAt, &task.UpdatedAt)
			if err != nil {
				log.Println("authorRepository.tasksRows.Next():", err.Error())
				return nil, err
			}
			author.Tasks = append(author.Tasks, task)
		}
		authors = append(authors, author)
	}
	return authors, nil
}

func (a *authorRepository) Get(id string) (model.Author, error) {
	var author model.Author
	err := a.db.QueryRow(config.SelectAuthorByID, id).Scan(
		&author.ID,
		&author.Name,
		&author.Email,
		&author.CreatedAt,
		&author.UpdatedAt,
	)
	if err != nil {
		log.Println("authorRepository.Get.QueryRow:", err.Error())
		return model.Author{}, err
	}
	rows, err := a.db.Query(config.SelectAuthorWithTask, id)
	if err != nil {
		log.Println("authorRepository.Get.Query:", err.Error())
		return model.Author{}, err
	}
	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Content, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			log.Println("authorRepository.rows.Next():", err.Error())
			return model.Author{}, err
		}
		author.Tasks = append(author.Tasks, task)
	}
	return author, nil
}

func NewAuthorRepository(db *sql.DB) AuthorRepository {
	return &authorRepository{db: db}
}
