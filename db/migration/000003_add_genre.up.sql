BEGIN;

CREATE TABLE genres(id BIGINT PRIMARY KEY, name TEXT);

CREATE TABLE books_genres(
    genre_id BIGINT REFERENCES genres(id),
    book_isbn TEXT REFERENCES books(isbn),
    PRIMARY KEY(genre_id, book_isbn)
);

END;