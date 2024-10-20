-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tender_objects (
    tender_id BIGINT NOT NULL REFERENCES tenders(id),
    object_id BIGINT NOT NULL REFERENCES objects(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tender_objects;
-- +goose StatementEnd
