ALTER TABLE projects
ADD CONSTRAINT unique_project_name UNIQUE (project_name);