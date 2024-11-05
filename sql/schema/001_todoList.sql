-- +goose Up
CREATE TABLE todolist (
  list TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  complete BOOLEAN NOT NULL

);

-- +goose Down
DROP TABLE todolist;   