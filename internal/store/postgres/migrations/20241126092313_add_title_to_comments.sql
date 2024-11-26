-- +goose Up
ALTER TABLE comments ADD COLUMN IF NOT EXISTS title TEXT;

UPDATE comments SET title = 'comment' WHERE title IS NULL;

ALTER TABLE comments ALTER COLUMN title SET NOT NULL;

-- +goose Down
ALTER TABLE comments DROP COLUMN IF EXISTS title;
