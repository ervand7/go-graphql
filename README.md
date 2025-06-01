# GraphQL API Guide

This guide provides instructions on how to use the GraphQL API for managing books and their authors.

---

## 🔧 Setup

1. **Install dependencies and generate GraphQL code**

   ```bash
   go mod tidy
   gqlgen generate
   ```

2. **Run the server**

   ```bash
   go run .
   ```

3. **Open the GraphQL Console**

   In your browser, go to: [http://localhost:8080](http://localhost:8080)

---

## 🚀 What This App Demonstrates

This project demonstrates how **GraphQL solves the underfetching problem**.

### ❌ Underfetching in REST:
To show a book and its author, REST clients often need to:
- First call `/books` → get `author_id`
- Then call `/authors/{id}` for each book → multiple round-trips

### ✅ In GraphQL:
You can make **a single query** like this:

```graphql
query {
  books {
    title
    author {
      name
    }
  }
}
```

And GraphQL will:
- Fetch all books
- Automatically call a resolver to fetch each `author`
- Combine it all into a single structured response

---

## 📚 GraphQL Queries and Mutations

### 🔍 Query: Fetch Books with Author Info

```graphql
query {
  books {
    id
    title
    author {
      id
      name
    }
  }
}
```

### ➕ Mutation: Add a Book (linked to existing author)

```graphql
mutation {
  addBook(title: "GraphQL in Action", authorID: "1") {
    id
    title
    author {
      name
    }
  }
}
```

---

## 📟 Using cURL

### 🔍 Fetch Books with Author Info

```bash
curl -X POST http://localhost:8080/graphql \
  -H "Content-Type: application/json" \
  -d '{"query": "query { books { id title author { id name } } }"}'
```

### ➕ Add a Book

```bash
curl -X POST http://localhost:8080/graphql \
  -H "Content-Type: application/json" \
  -d '{"query": "mutation { addBook(title: \"GraphQL in Action\", authorID: \"1\") { id title author { name } } }"}'
```

---

## 📌 Notes

- Authors are predefined in-memory (e.g., `"1"` = Robert C. Martin)
- The `author` field is resolved separately using `@goField(forceResolver: true)`
- This approach demonstrates **how GraphQL handles nested data fetching seamlessly**

---

Happy querying! 🎯
