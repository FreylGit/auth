-- +goose Up
CREATE TABLE refresh_tokens(
    id SERIAL PRIMARY KEY ,
    token BYTEA,
    exp timestamp NOT NULL
);
CREATE INDEX idx_token ON refresh_tokens (token);

-- +goose Down
DROP INDEX idx_token;
DROP TABLE refresh_tokens;
