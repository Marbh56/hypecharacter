-- +goose Up

ALTER TABLE users
    ADD COLUMN email TEXT,
    ADD COLUMN password_hash TEXT;

-- +goose Down

ALTER TABLE users
    DROP COLUMN email,
    DROP COLUMN password_hash;
