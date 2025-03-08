CREATE TABLE projects (
    id UUID PRIMARY KEY,
    project_name VARCHAR(255) NOT NULL,
    channel_id VARCHAR(255) NOT NULL,
    added_by VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    description TEXT
);