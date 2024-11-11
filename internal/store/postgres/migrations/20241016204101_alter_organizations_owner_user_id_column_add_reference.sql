-- +goose Up
ALTER TABLE IF EXISTS organizations ADD FOREIGN KEY (owner_user_id) REFERENCES users(id);

-- +goose Down
ALTER TABLE IF EXISTS organizations DROP CONSTRAINT organizations_owner_user_id_fkey;
