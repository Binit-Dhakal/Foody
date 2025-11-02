ALTER TABLE users ADD COLUMN last_login TIMESTAMP;

CREATE TABLE sessions (
    session_key TEXT PRIMARY KEY,
    user_id BIGINT REFERENCES users(id) ON DELETE CASCADE,
    token_hash BYTEA NOT NULL,
    scope TEXT,
    expires_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

