package graph

import (
	"context"
	"fmt"

	"go-graphql/graph/generated"
	"go-graphql/graph/model"
)

func (r *Resolver) Query() generated.QueryResolver       { return &queryResolver{r} }
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type queryResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }

func (q *queryResolver) Books(ctx context.Context) ([]*model.Book, error) {
	out := make([]*model.Book, len(q.books))
	for i, b := range q.books {
		out[i] = &model.Book{ID: b.ID, Title: b.Title, Author: b.Author}
	}
	return out, nil
}

func (m *mutationResolver) AddBook(ctx context.Context, title string, author string) (*model.Book, error) {
	id := fmt.Sprintf("%d", len(m.books)+1)
	book := Book{ID: id, Title: title, Author: author}
	m.books = append(m.books, book)

	return &model.Book{ID: book.ID, Title: book.Title, Author: book.Author}, nil
}
