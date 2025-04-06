-- Write your 'down' migration SQL here
ALTER TABLE projects
DROP COLUMN project_source_auth_token;