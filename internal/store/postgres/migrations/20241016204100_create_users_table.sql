-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id              BIGSERIAL PRIMARY KEY,
    organization_id BIGINT NOT NULL REFERENCES organizations(id),
    email           TEXT NOT NULL UNIQUE,
    phone           TEXT NOT NULL,
    password_hash   TEXT NOT NULL,
    totp_salt       TEXT NOT NULL,
    first_name      TEXT NOT NULL,
    last_name       TEXT NOT NULL,
    middle_name     TEXT NOT NULL,
    avatar_url      TEXT NULL,
    email_verified  BOOLEAN DEFAULT FALSE,
    role            SMALLINT DEFAULT 1,
    created_at      TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS users;
