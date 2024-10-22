-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tender_responses (
    tender_id       BIGINT  NOT NULL REFERENCES tenders(id),
    organization_id BIGINT  NOT NULL REFERENCES organizations(id),
    price           INT     NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tender_responses;
-- +goose StatementEnd
