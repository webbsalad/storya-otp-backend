-- +goose Up
-- +goose StatementBegin

CREATE TABLE otp_keys
(
    id UUID DEFAULT uuid_generate_v4() NOT NULL CONSTRAINT otp_keys_pkey PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    otp_key TEXT NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE otp_keys;

-- +goose StatementEnd
