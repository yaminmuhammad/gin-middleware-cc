package config

const (
	SelectTaskList       = `SELECT id, title, content, created_at, updated_at FROM tasks ORDER BY created_at desc`
	InsertTask           = `INSERT INTO tasks (title, content, author_id, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	SelectTaskByAuthorID = `SELECT id, title, content, author_id, created_at, updated_at FROM tasks WHERE id = $1 ORDER BY created_at DESC`

	SelectAuthorByID = `SELECT id, name, email, created_at, updated_at FROM authors WHERE id = $1`
	// SelectAuthorList = `SELECT ID, name, email, created_at, updated_at FROM authors ORDER BY created_at desc`
	// InsertAuthor     = `INSERT INTO authors (name, email, updated_at) VALUES($1, $2, $3) RETURNING id, created_at`
)
