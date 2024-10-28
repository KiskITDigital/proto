-- +goose Up
CREATE SCHEMA IF NOT EXISTS admin;

-- +goose Down
DROP SCHEMA IF EXISTS admin;
