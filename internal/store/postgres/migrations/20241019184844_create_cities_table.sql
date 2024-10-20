-- +goose Up
-- +goose StatementBegin
CREATE TABLE cities (
    id          BIGSERIAL PRIMARY KEY,
    region_id   BIGINT NOT NULL REFERENCES regions(id),
    name        TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cities;
-- +goose StatementEnd
