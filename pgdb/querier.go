// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package pgdb

import (
	"context"
)

type Querier interface {
	CreateAuthor(ctx context.Context, arg CreateAuthorParams) (Author, error)
	CreateBook(ctx context.Context, arg CreateBookParams) (Book, error)
	DeleteAuthor(ctx context.Context, id int64) error
	DeleteBook(ctx context.Context, id int64) error
	GetAuthor(ctx context.Context, id int64) (Author, error)
	GetBook(ctx context.Context, id int64) (Book, error)
	GetBookWithAuthor(ctx context.Context) (GetBookWithAuthorRow, error)
	ListAuthors(ctx context.Context) ([]Author, error)
	ListBooks(ctx context.Context) ([]Book, error)
	UpdateAuthor(ctx context.Context, arg UpdateAuthorParams) error
	UpdateBook(ctx context.Context, arg UpdateBookParams) error
}

var _ Querier = (*Queries)(nil)
