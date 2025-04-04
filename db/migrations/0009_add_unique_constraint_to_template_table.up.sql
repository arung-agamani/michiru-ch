-- Write your 'up' migration SQL here
ALTER TABLE templates
ADD CONSTRAINT unique_project_event_type UNIQUE (project_id, event_type);