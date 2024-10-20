-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tender_services (
    tender_id BIGINT NOT NULL REFERENCES tenders(id),
    service_id BIGINT NOT NULL REFERENCES services(id)
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tender_services;
-- +goose StatementEnd
