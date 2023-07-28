CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- users table
CREATE TABLE IF NOT EXISTS users (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name varchar(128) NOT NULL,
    email varchar(256) UNIQUE NOT NULL,
    password varchar(256) NOT NULL
);

-- tags table
CREATE TABLE IF NOT EXISTS tags (
    id bigserial PRIMARY KEY,
    name varchar(64) UNIQUE NOT NULL
);

-- posts table
CREATE TABLE IF NOT EXISTS posts (
    id bigserial PRIMARY KEY,
    title varchar(256) NOT NULL,
    body varchar NOT NULL,
    user_id uuid NOT NULL,
    tag_ids bigint[] NOT NULL,
    created_at timestamptz NOT NULL DEFAULT(NOW()),
    updated_at timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z')
);

ALTER TABLE
    "posts"
ADD
    FOREIGN KEY ("user_id") REFERENCES "users" ("id");

-- comments table
CREATE TABLE IF NOT EXISTS comments (
    id bigserial PRIMARY KEY,
    body varchar NOT NULL,
    post_id bigint NOT NULL,
    created_at timestamptz NOT NULL DEFAULT(NOW()),
    updated_at timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z')
);

ALTER TABLE
    "comments"
ADD
    FOREIGN KEY ("post_id") REFERENCES "posts" ("id");