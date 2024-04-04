-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
                       id INTEGER PRIMARY KEY,
                       username TEXT NOT NULL,
                       email TEXT NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
