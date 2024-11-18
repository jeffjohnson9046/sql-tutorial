-- name: GetBook :one
SELECT
    * 
FROM
    books
WHERE
    id = $1
LIMIT 1;

-- name: GetBookWithAuthor :one
SELECT
    a.name,
    b.title,
    b.isbn,
    b.subject
FROM
    books b
JOIN
    authors a
    ON (a.id = b.author_id)
ORDER BY
    a.name,
    b.title;

-- name: ListBooks :many
SELECT
    *
FROM
    books
ORDER BY
    title;

-- name: CreateBook :one
INSERT INTO books (author_id, title, isbn, subject)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: UpdateBook :exec
UPDATE 
    books
SET 
    title = $2,
    isbn = $3,
    subject = $4
WHERE
    id = $1;

-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1;