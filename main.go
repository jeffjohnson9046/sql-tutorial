package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"tutorial.sqlc.dev/app/pgdb"
)

func run() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "user=jeffjohnson dbname=sqlc_test sslmode=verify-full")
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	db := pgdb.New(conn)

	// list all authors
	authors, err := db.ListAuthors(ctx)
	if err != nil {
		return err
	}
	log.Println(authors)

	// create an author
	insertedAuthor, err := db.CreateAuthor(ctx,
		pgdb.CreateAuthorParams{
			Name: "Brian Kernighan",
			Bio:  pgtype.Text{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
		})
	if err != nil {
		return err
	}
	log.Println(insertedAuthor)

	// get the author we just inserted
	author, err := db.GetAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		return err
	}
	log.Println(author)

	// update the author
	foo := db.UpdateAuthor(ctx, pgdb.UpdateAuthorParams{Name: "Darth Vader", Bio: pgtype.Text{String: "Dark Lord of the Sith", Valid: true}, ID: author.ID})
	if foo != nil {
		return foo
	}

	updatedAuthor, err := db.GetAuthor(ctx, author.ID)
	if err != nil {
		return err
	}
	log.Println(updatedAuthor)

	newBook, err := db.CreateBook(ctx, pgdb.CreateBookParams{Title: "How to be a Dark Lord in 12 Easy Steps", AuthorID: updatedAuthor.ID, Isbn: "some value"})
	if err != nil {
		return err
	}
	log.Println(newBook)

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
