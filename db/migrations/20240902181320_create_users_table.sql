-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
     user_id SERIAL PRIMARY KEY,
     firstname VARCHAR(100) NOT NULL,
     lastname VARCHAR(100) NOT NULL,
     phone_number VARCHAR(15) NOT NULL UNIQUE,
     email VARCHAR(100) NOT NULL UNIQUE,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
 );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
