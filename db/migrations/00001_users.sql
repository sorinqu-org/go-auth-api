-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
