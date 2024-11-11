-- +goose Up
CREATE TABLE IF NOT EXISTS tender_responses (
    tender_id       BIGINT      NOT NULL REFERENCES tenders(id),
    organization_id BIGINT      NOT NULL REFERENCES organizations(id),
    price           INT         NOT NULL,
    is_nds_price    BOOLEAN     NOT NULL,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS tender_responses;
