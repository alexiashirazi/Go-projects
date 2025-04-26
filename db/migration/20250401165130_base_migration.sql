-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name text not null,
    last_name text not null,
    email text UNIQUE not null,
    password text not null
);

-- +goose Down
DROP TABLE IF EXISTS users;

