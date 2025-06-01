package graph

var (
	authors = map[string]Author{
		"1": {ID: "1", Name: "Robert C. Martin"},
		"2": {ID: "2", Name: "Alan Donovan"},
	}

	books = []Book{
		{ID: "1", Title: "Clean Code", AuthorID: "1"},
		{ID: "2", Title: "The Go Programming Language", AuthorID: "2"},
	}
)

type Resolver struct {
	books   []Book
	authors map[string]Author
}

type Author struct {
	ID   string
	Name string
}

type Book struct {
	ID       string
	Title    string
	AuthorID string
}

func NewResolver() *Resolver {

	return &Resolver{
		books:   books,
		authors: authors,
	}
}
