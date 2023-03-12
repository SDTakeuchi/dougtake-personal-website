CREATE TABLE IF NOT EXISTS sessions (
    id uuid PRIMARY KEY,
    refresh_token varchar NOT NULL,
    user_id uuid NOT NULL,
    user_agent varchar,
    client_ip varchar,
    expires_at timestamptz NOT NULL,
    created_at timestamptz NOT NULL DEFAULT(NOW())
);

ALTER TABLE
    "sessions"
ADD
    FOREIGN KEY ("user_id") REFERENCES "user" ("id");
