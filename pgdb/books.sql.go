// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: books.sql

package pgdb

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createBook = `-- name: CreateBook :one
INSERT INTO books (author_id, title, isbn, subject)
VALUES (
    $1,
    $2,
    $3,
    $4
)
RETURNING id, author_id, title, isbn, subject
`

type CreateBookParams struct {
	AuthorID int64       `json:"authorId"`
	Title    string      `json:"title"`
	Isbn     string      `json:"isbn"`
	Subject  pgtype.Text `json:"subject"`
}

func (q *Queries) CreateBook(ctx context.Context, arg CreateBookParams) (Book, error) {
	row := q.db.QueryRow(ctx, createBook,
		arg.AuthorID,
		arg.Title,
		arg.Isbn,
		arg.Subject,
	)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.AuthorID,
		&i.Title,
		&i.Isbn,
		&i.Subject,
	)
	return i, err
}

const deleteBook = `-- name: DeleteBook :exec
DELETE FROM books
WHERE id = $1
`

func (q *Queries) DeleteBook(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteBook, id)
	return err
}

const getBook = `-- name: GetBook :one
SELECT
    id, author_id, title, isbn, subject 
FROM
    books
WHERE
    id = $1
LIMIT 1
`

func (q *Queries) GetBook(ctx context.Context, id int64) (Book, error) {
	row := q.db.QueryRow(ctx, getBook, id)
	var i Book
	err := row.Scan(
		&i.ID,
		&i.AuthorID,
		&i.Title,
		&i.Isbn,
		&i.Subject,
	)
	return i, err
}

const getBookWithAuthor = `-- name: GetBookWithAuthor :one
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
    b.title
`

type GetBookWithAuthorRow struct {
	Name    string      `json:"name"`
	Title   string      `json:"title"`
	Isbn    string      `json:"isbn"`
	Subject pgtype.Text `json:"subject"`
}

func (q *Queries) GetBookWithAuthor(ctx context.Context) (GetBookWithAuthorRow, error) {
	row := q.db.QueryRow(ctx, getBookWithAuthor)
	var i GetBookWithAuthorRow
	err := row.Scan(
		&i.Name,
		&i.Title,
		&i.Isbn,
		&i.Subject,
	)
	return i, err
}

const listBooks = `-- name: ListBooks :many
SELECT
    id, author_id, title, isbn, subject
FROM
    books
ORDER BY
    title
`

func (q *Queries) ListBooks(ctx context.Context) ([]Book, error) {
	rows, err := q.db.Query(ctx, listBooks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Book{}
	for rows.Next() {
		var i Book
		if err := rows.Scan(
			&i.ID,
			&i.AuthorID,
			&i.Title,
			&i.Isbn,
			&i.Subject,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateBook = `-- name: UpdateBook :exec
UPDATE 
    books
SET 
    title = $2,
    isbn = $3,
    subject = $4
WHERE
    id = $1
`

type UpdateBookParams struct {
	ID      int64       `json:"id"`
	Title   string      `json:"title"`
	Isbn    string      `json:"isbn"`
	Subject pgtype.Text `json:"subject"`
}

func (q *Queries) UpdateBook(ctx context.Context, arg UpdateBookParams) error {
	_, err := q.db.Exec(ctx, updateBook,
		arg.ID,
		arg.Title,
		arg.Isbn,
		arg.Subject,
	)
	return err
}