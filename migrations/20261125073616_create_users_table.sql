-- +goose Up
CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE users (
    id            uuid PRIMARY KEY DEFAULT uuidv7(),
    email         citext NOT NULL UNIQUE,
    password_hash text NOT NULL,
    full_name     text NOT NULL,
    platform_role text NOT NULL DEFAULT 'user' CHECK (platform_role IN ('user', 'super_admin')),
    status        text NOT NULL DEFAULT 'active' CHECK (status IN ('active', 'disabled')),
    created_at    timestamptz NOT NULL DEFAULT now(),
    updated_at    timestamptz NOT NULL DEFAULT now()
);

-- +goose Down
DROP TABLE IF EXISTS users;