package repository

import (
	"database/sql"
	"gin/config"
	"gin/model"
	"log"
)

// TODO:
/**
1. Interface (v) -> kontrak (CRUD), mudahkan unit testing
3. Struct (v) -> data acces (db)
2. Method (v) -> implementasi dari interface
3. Function (v) -> constructor (gerbang penghubung)
*/
type TaskRepository interface {
	List() ([]model.Task, error)
	Create(payload model.Task) (model.Task, error)
	GetByAuthor(authorId string) ([]model.Task, error)
}

type taskRepository struct {
	db *sql.DB
}

func (t *taskRepository) GetByAuthor(authorId string) ([]model.Task, error) {
	var tasks []model.Task
	rows, err := t.db.Query(config.SelectTaskByAuthorID, authorId)
	if err != nil {
		log.Println("taskRepository.GetByAuthor.Query:", err.Error())
		return nil, err
	}
	for rows.Next() {
		var task model.Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Content,
			&task.AuthorId,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			log.Println("taskRepository.GetByAuthor.Rows:", err.Error())
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (t *taskRepository) List() ([]model.Task, error) {
	var tasks []model.Task
	rows, err := t.db.Query(config.SelectTaskList)
	if err != nil {
		log.Println("taskRepository.Query:", err.Error())
		return nil, err
	}
	for rows.Next() {
		var task model.Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Content,
			&task.AuthorId,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			log.Println("taskRepository.Rows.Next():", err.Error())
			return nil, err
		}

		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (t *taskRepository) Create(payload model.Task) (model.Task, error) {
	var task model.Task
	err := t.db.QueryRow(config.InsertTask,
		payload.Title,
		payload.Content,
		payload.AuthorId,
		payload.UpdatedAt).Scan(
		&task.ID,
		&task.CreatedAt,
	)
	if err != nil {
		log.Println("taskRepository.QueryRow:", err.Error())
		return model.Task{}, err
	}
	task.Title = payload.Title
	task.Content = payload.Content
	task.AuthorId = payload.AuthorId
	task.UpdatedAt = payload.UpdatedAt
	return task, nil
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{db: db}
}
