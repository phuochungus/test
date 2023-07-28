CREATE TABLE IF NOT EXISTS
    authors (
        id BIGSERIAL PRIMARY KEY,
        name text NOT NULL,
        bio text,
        date_of_birth TIMESTAMP NOT NULL
    );