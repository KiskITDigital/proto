-- +goose Up
-- +goose StatementBegin
CREATE TABLE regions (
    id BIGSERIAL PRIMARY KEY,
    name TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS regions;
-- +goose StatementEnd
