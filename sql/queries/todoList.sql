-- name: CreateTodo :one
INSERT INTO todolist(id, list, created_at, updated_at, complete)
VALUES($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetTodoList :many
SELECT * FROM todolist;



-- name: CompleteToDo :one
UPDATE todolist
SET complete = TRUE
WHERE id = $1
RETURNING *;