ALTER TABLE projects ADD COLUMN webhook_url VARCHAR(255);
ALTER TABLE projects ADD COLUMN webhook_secret VARCHAR(255);
ALTER TABLE projects ADD COLUMN webhook_origin VARCHAR(255);