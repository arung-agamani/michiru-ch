ALTER TABLE projects
DROP COLUMN webhook_url;

ALTER TABLE projects
DROP COLUMN webhook_secret;

ALTER TABLE projects
DROP COLUMN webhook_origin;