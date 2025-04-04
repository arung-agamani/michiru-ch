-- Write your 'down' migration SQL here
ALTER TABLE templates
DROP CONSTRAINT unique_project_event_type;