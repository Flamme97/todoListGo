-- +goose Up
CREATE TABLE todolist (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  list TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  complete BOOLEAN NOT NULL DEFAULT FALSE

);

-- +goose Down
DROP TABLE todolist;   