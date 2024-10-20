-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tenders (
    id                  BIGSERIAL       PRIMARY KEY,
    name                TEXT            NOT NULL,
    price               INT             NOT NULL,
    is_contract_price   BOOLEAN         NOT NULL DEFAULT FALSE,
    is_nds_price        BOOLEAN         NOT NULL DEFAULT FALSE,
    city_id             INT             NOT NULL REFERENCES cities(id),
    floor_space         INT             NOT NULL,
    description         TEXT            NULL,
    wishes              TEXT            NULL,
    specification       TEXT            NOT NULL,
    attachments         TEXT[]          NULL,
    verified            BOOLEAN         NOT NULL DEFAULT FALSE,
    is_draft            BOOLEAN         NOT NULL,
    reception_start     TIMESTAMPTZ     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    reception_end       TIMESTAMPTZ     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    work_start          TIMESTAMPTZ     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    work_end            TIMESTAMPTZ     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    organization_id     BIGINT          NOT NULL REFERENCES organizations(id),
    created_at          TIMESTAMPTZ     DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tenders;
-- +goose StatementEnd
