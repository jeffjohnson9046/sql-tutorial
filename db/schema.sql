CREATE TABLE IF NOT EXISTS authors (
    id bigserial PRIMARY KEY,
    name text NOT NULL,
    bio text
);

CREATE TABLE IF NOT EXISTS books (
    id bigserial PRIMARY KEY,
    author_id bigint NOT NULL REFERENCES authors(id),
    title text NOT NULL,
    isbn text NOT NULL,
    subject text
);