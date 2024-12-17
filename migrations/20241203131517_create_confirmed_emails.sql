-- +goose Up
-- +goose StatementBegin

CREATE TABLE confirmed_emails
(
    id UUID DEFAULT uuid_generate_v4() NOT NULL CONSTRAINT confirmed_emails_pkey PRIMARY KEY,
    email TEXT NOT NULL UNIQUE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE confirmed_emails;

-- +goose StatementEnd
