package graph

type Resolver struct {
	books []Book
}

type Book struct {
	ID, Title, Author string
}

var books = []Book{
	{ID: "1", Title: "Clean Code", Author: "Robert C. Martin"},
	{ID: "2", Title: "The Go Programming Language", Author: "Alan Donovan"},
}

func NewResolver() *Resolver {
	return &Resolver{books: books}
}
