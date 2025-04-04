-- Write your 'up' migration SQL here
-- Create the templates table
CREATE TABLE templates (
    id SERIAL PRIMARY KEY,
    project_id UUID NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    event_type VARCHAR(255) NOT NULL,
    template TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create the predefined_templates table
CREATE TABLE predefined_templates (
    id SERIAL PRIMARY KEY,
    event_type VARCHAR(255) NOT NULL,
    template TEXT NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);