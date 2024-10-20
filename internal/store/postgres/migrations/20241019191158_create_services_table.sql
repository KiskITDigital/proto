-- +goose Up
-- +goose StatementBegin
CREATE TABLE services (
    id          BIGSERIAL PRIMARY KEY,
    parent_id   BIGINT REFERENCES services(id) ON DELETE CASCADE,
    name        TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS services;
-- +goose StatementEnd
