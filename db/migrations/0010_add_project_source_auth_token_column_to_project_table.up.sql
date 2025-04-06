-- Write your 'up' migration SQL here
ALTER TABLE projects
ADD COLUMN project_source_auth_token VARCHAR(255) DEFAULT '';