# GraphQL API Guide

This guide provides instructions on how to use the GraphQL API for managing books.

## Setup

1. **Generate GraphQL Code**
   
   Run the following command to generate the necessary GraphQL code:
   
   ```bash
   gqlgen generate
   ```

2. **Access the GraphQL Console**

   Open your browser and go to the following URL to access the GraphQL console:
   
   [http://localhost:8080](http://localhost:8080)

## Queries and Mutations

### Query: Fetch Books

To fetch the list of books, run the following query:

```graphql
query {
  books {
    id
    title
    author
  }
}
```

### Mutation: Add a Book

To add a new book, use the following mutation:

```graphql
mutation {
  addBook(title: "GraphQL in Action", author: "Samer Buna") {
    id
    title
    author
  }
}
```

## Using cURL

You can also perform the above operations using cURL commands:

### Fetch Books

```bash
curl -X POST http://localhost:8080/graphql \
  -H "Content-Type: application/json" \
  -d '{"query": "query { books { id title author } }"}'
```

### Add a Book

```bash
curl -X POST http://localhost:8080/graphql \
  -H "Content-Type: application/json" \
  -d '{"query": "mutation { addBook(title: \\"GraphQL in Action\\", author: \\"Samer Buna\\") { id title author } }"}'
``` 