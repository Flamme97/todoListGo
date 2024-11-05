-- name: CreateTodo :one
INSERT INTO todolist(list, created_at, updated_at)
VALUES($1, $2, $3)
RETURNING *;

-- name: GetTodoList :many
SELECT * FROM todolist;

