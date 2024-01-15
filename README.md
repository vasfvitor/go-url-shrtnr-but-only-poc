Hands-On RESTful Web Services with Go - Reader's challenge – an API for URL shortening

# Task

1. Under the hood, the following things happen in a URL shortening service:
Take the original URL

2. Apply BASE62 encoding on it; it generates a Shortened URL

3. Store that URL in the database. Map it to the original URL ([shortened_url:original_url])

4. Whenever a request comes to the shortened URL, just do an HTTP redirect to the original URL

Node: You can use a dummy JSON file/Go map to store the URL for now instead of a database.


# Endpoints 

| URL                   | REST Verb | Action                              | Success | Failure  |
| --------------------- | --------- | ----------------------------------- | ------- | -------- |
| /api/v1/new           | POST      | Create a shortened URL              | 200     | 500, 404 |
| /api/v1/:url          | GET       | Redirect to original URL            | 301     | 404      |
| /api/v1/debug/listall | GET       | List all maps to the server console |         |          |


# Implementation:

gorilla/mux

# Takeaways:

- Struct and other types can be inlined
- I need to learn about un/marshalling and its relation to de/serialization
- This could be done with std lib as of Go 1.22

NTS: Don't lose time in stuff outside the exercise scope, in this case, tests and design choices.