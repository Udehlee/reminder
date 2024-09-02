-- +goose Up
-- +goose StatementBegin
CREATE TABLE reminders (
    reminder_id SERIAL PRIMARY KEY,
    user_reminder INT REFERENCES users(user_id) ON DELETE CASCADE,
    sent_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    message TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
