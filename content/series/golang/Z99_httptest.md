# `httptest`

The [`httptest`](https://golang.org/pkg/net/http/httptest/) package provides utilities for HTTP testing. You can use `httptest` to write end-to-end (E2E) tests of your API endpoints, for example. `httptest` is part of the `net/http` module in the standard library, and so you don't need to import any third-party tools.

As with any type of HTTP testing, the process includes 3 steps:

1. Building the request
2. Sending the request and capturing the response
3. Making assertions on:
   - the response
   - changes on external states (e.g. changes to database and/or filesystem)

But typically, you'd need to actually spin up a test instance of your server

## Building the Request

