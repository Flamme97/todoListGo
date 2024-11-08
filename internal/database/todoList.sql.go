// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: todoList.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const completeToDo = `-- name: CompleteToDo :one
UPDATE todolist
SET complete = TRUE
WHERE id = $1
RETURNING id, list, created_at, updated_at, complete
`

func (q *Queries) CompleteToDo(ctx context.Context, id uuid.UUID) (Todolist, error) {
	row := q.db.QueryRowContext(ctx, completeToDo, id)
	var i Todolist
	err := row.Scan(
		&i.ID,
		&i.List,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Complete,
	)
	return i, err
}

const createTodo = `-- name: CreateTodo :one
INSERT INTO todolist(id, list, created_at, updated_at, complete)
VALUES($1, $2, $3, $4, $5)
RETURNING id, list, created_at, updated_at, complete
`

type CreateTodoParams struct {
	ID        uuid.UUID
	List      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Complete  bool
}

func (q *Queries) CreateTodo(ctx context.Context, arg CreateTodoParams) (Todolist, error) {
	row := q.db.QueryRowContext(ctx, createTodo,
		arg.ID,
		arg.List,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Complete,
	)
	var i Todolist
	err := row.Scan(
		&i.ID,
		&i.List,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Complete,
	)
	return i, err
}

const getTodoList = `-- name: GetTodoList :many
SELECT id, list, created_at, updated_at, complete FROM todolist
`

func (q *Queries) GetTodoList(ctx context.Context) ([]Todolist, error) {
	rows, err := q.db.QueryContext(ctx, getTodoList)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todolist
	for rows.Next() {
		var i Todolist
		if err := rows.Scan(
			&i.ID,
			&i.List,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Complete,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
