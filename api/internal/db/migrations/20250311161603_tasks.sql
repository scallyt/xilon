-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    task_name TEXT NOT NULL,
    todo TEXT,
    status TEXT CHECK (status IN ('in progress', 'completed', 'pending')) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE tasks;
-- +goose StatementEnd
