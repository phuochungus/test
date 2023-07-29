CREATE TABLE authors (
    id BIGINT PRIMARY KEY,
    name TEXT NOT NULL,
    bio TEXT,
    date_of_birth TIMESTAMP NOT NULL
);