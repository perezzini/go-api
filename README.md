# Dummy API in Golang language

[![CI](https://github.com/perezzini/go-api/actions/workflows/integrate.yml/badge.svg)](https://github.com/perezzini/go-api/actions/workflows/integrate.yml)

Just for exercising Golang fundaments.

API implemented using the [Gin](https://github.com/gin-gonic/gin) framework.

> This is a WIP project

## Endpoints

### Books

- `GET /books`: Retrieve all books
- `POST /books`: Post a new book
- `GET /books/{id}: Retrive book with ID `id`
- `PATCH /books/{id}/checkout`: Checkout book with ID `id`
- `PATCH /books/{id}/return`: Return book with ID `id`
