-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS sessions(
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    session VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    timeout_at TIMESTAMP DEFAULT (CURRENT_TIMESTAMP + INTERVAL 1 HOUR),
    UNIQUE INDEX sessions_session_index (session)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS sessions CASCADE;
-- +goose StatementEnd