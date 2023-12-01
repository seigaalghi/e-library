# e-library
## How to run
- Go v1.18.1 or newer
- Run `go mod tidy`
- Run `go run main.go`

## Stack Used
- Sqlite3 DB
- Mockery
- Go
- Zap Logger
- Go-Kit
- Mux Router
- JWT

## APIs
### Login
- Use any random email and password
- The login api is just mocked
- But the token still required

```
curl --location 'localhost:8000/api/v1/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email" : "some@gmail.com",
    "password" : "somepw"
}'
```

### Get Books By Subject
- Use the generated token from login proccess for the get books request

```
curl --location 'localhost:8000/api/v1/books?subject=love&page=1&limit=50' \
--header 'Authorization: Bearer Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNvbWVAZ21haWwuY29tIiwiZXhwIjoxNzAxNDAxNDk0LCJpYXQiOjE3MDE0MDA1OTR9.4IhB0A5sGvQc611ykM5MITRx0iDY_fVQTFZK-bYMNpE'
```

### Lend Book
- Pick one edition number from get books request
- Add `pickup_date` and `dropoff_date` to the request body
- The desired date format `yyyy-mm-dd`
- The book cannot be borrowed again until the book dropped off by the user.

```
curl --location 'localhost:8000/api/v1/books/lend' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNvbWVAZ21haWwuY29tIiwiZXhwIjoxNzAxNDAxNDk0LCJpYXQiOjE3MDE0MDA1OTR9.4IhB0A5sGvQc611ykM5MITRx0iDY_fVQTFZK-bYMNpE' \
--data '{
    "edition_number" : "OL38586477M",
    "pickup_date" : "2023-12-10",
    "dropoff_date" : "2023-12-14"
}'
```