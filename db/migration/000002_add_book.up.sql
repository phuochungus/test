BEGIN;

CREATE TABLE publishers (id BIGINT PRIMARY KEY, name TEXT);

CREATE TABLE books (
    isbn TEXT PRIMARY KEY,
    name TEXT,
    publisher_id BIGINT REFERENCES publishers (id),
    quantity INT DEFAULT 0,
    published_year SMALLINT
);

CREATE TABLE books_authors (
    book_isbn TEXT REFERENCES books(isbn),
    authors_id BIGINT REFERENCES authors(id),
    PRIMARY KEY (book_isbn, authors_id)
);

COMMIT;