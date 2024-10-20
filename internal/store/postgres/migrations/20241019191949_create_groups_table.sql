-- +goose Up
-- +goose StatementBegin
CREATE TABLE objects (
    id          BIGSERIAL PRIMARY KEY,
    parent_id   BIGINT REFERENCES objects(id) ON DELETE CASCADE,
    name        TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS objects;
-- +goose StatementEnd
