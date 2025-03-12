-- name: CreateProject :one
INSERT INTO projects (customer_name, customer_id, project_name, chat_id, description, developer_id, status, start_date, end_date, budget)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: GetProject :one
SELECT *
FROM projects
WHERE id = $1;

-- name: GetProjects :many
SELECT *
FROM projects
ORDER BY start_date DESC
LIMIT $1
OFFSET $2;

-- name: UpdateProject :one
UPDATE projects
SET customer_name = $2,
  customer_id = $3,
  project_name = $4,
  chat_id = $5,
  description = $6,
  developer_id = $7,
  status = $8,
  start_date = $9,
  end_date = $10,
  budget = $11
WHERE id = $1
RETURNING *;

-- name: DeleteProject :exec
DELETE FROM projects
WHERE id = $1;

-- name: GetAllProjectByUserId :many
SELECT *
FROM projects
WHERE developer_id = $1
ORDER BY start_date DESC;

-- name: GetOneProjectById :one
SELECT *
FROM projects
WHERE id = $1;

-- name: GetProjectsByStatus :many
SELECT *
FROM projects
WHERE status = $1
ORDER BY start_date DESC
LIMIT $2
OFFSET $3;