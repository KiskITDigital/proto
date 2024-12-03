-- +goose Up
CREATE TABLE IF NOT EXISTS comments (
    id                  BIGSERIAL PRIMARY KEY,
    organization_id     BIGINT REFERENCES organizations(id),
    title               TEXT NOT NULL,
    object_type         SMALLINT NOT NULL,
    object_id           BIGINT NOT NULL,
    content             TEXT NOT NULL,
    attachments         TEXT[],
    verification_status SMALLINT NOT NULL,
    created_at          TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS comments;