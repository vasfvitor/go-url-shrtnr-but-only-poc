Hands-On RESTful Web Services with Go - Reader's challenge – an API for URL
shortening

# Task

Under the hood, the following things happen in a URL shortening service:
Take the original URL1.
Apply BASE62 encoding on it; it generates a Shortened URL2.
Store that URL in the database. Map it to the original URL ([shortened_url:3.
original_url])
Whenever a request comes to the shortened URL, just do an HTTP redirect to the4.
original URL

Node: You can use a dummy JSON file/Go map to store the URL for now instead
of a database.


# Endpoints 

| URL                   | REST Verb | Action                              | Success | Failure  |
| --------------------- | --------- | ----------------------------------- | ------- | -------- |
| /api/v1/new           | POST      | Create a shortened URL              | 200     | 500, 404 |
| /api/v1/:url          | GET       | Redirect to original URL            | 301     | 404      |
| /api/v1/debug/listall | GET       | List all URLs to the server console |         |          |


# Implementation:

gorilla/mux