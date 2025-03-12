-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TABLE projects (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    customer_name TEXT NOT NULL,
    customer_id UUID NOT NULL,
    project_name TEXT NOT NULL,
    chat_id UUID UNIQUE,
    description TEXT,
    developer_id UUID NOT NULL,
    status TEXT CHECK (status IN ('in progress', 'completed', 'pending')) NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE,
    budget NUMERIC(12, 2) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE projects;
-- +goose StatementEnd
