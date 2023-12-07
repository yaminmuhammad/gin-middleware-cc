package config

const (
	SelectTaskList   = `SELECT id, title, content, created_at, updated_at FROM tasks ORDER BY created_at desc`
	SelectAuthorList = `
	SELECT
    a.id,
    a.name,
    a.email,
    a.created_at,
    a.updated_at,
    t.id AS task_id,
    t.title AS task_title,
    t.content AS task_content,
    t.created_at AS task_created_at,
    t.updated_at AS task_updated_at
	FROM
    authors a
	LEFT JOIN
    tasks t ON a.id = t.author_id
	ORDER BY
    a.created_at DESC
`
	InsertTask           = `INSERT INTO tasks (title, content, author_id, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	SelectTaskByAuthorID = `SELECT id, title, content, author_id, created_at, updated_at FROM tasks WHERE author_id = $1 ORDER BY created_at DESC`
	SelectAuthorByID     = `SELECT id, name, email, created_at, updated_at FROM authors WHERE id = $1`
)
