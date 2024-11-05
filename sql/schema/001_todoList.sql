-- +goose Up
CREATE TABLE todolist (
  list TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  complete BOOLEAN

);

-- +goose Down
DROP TABLE todolist;   