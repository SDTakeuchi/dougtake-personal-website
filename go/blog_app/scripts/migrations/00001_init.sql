CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id uuid DEFAULT uuid_generate_v4(),
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
);

CREATE TABLE IF NOT EXISTS tags (
    id bigserial,
    name varchar(255) NOT NULL,
);

CREATE TABLE IF NOT EXISTS posts (
    id bigserial,
    title varchar(255) NOT NULL,
    body varchar(255) NOT NULL,
    user_id uuid NOT NULL,
    tag_ids bigint(20)[] NOT NULL,
    created_at datetime NOT NULL,
    updated_at datetime NOT NULL
);

CREATE TABLE IF NOT EXISTS comments (
    id bigserial,
    body varchar(255) NOT NULL,
    post_id bigint(20) NOT NULL,
    created_at datetime NOT NULL,
    updated_at datetime NOT NULL,
);
