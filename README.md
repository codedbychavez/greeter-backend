# greeter-backend
A Golang backend for greeter


### Start

1. go run main.go

### How it works

1. App exposes API Endpoint: http://localhost:4000/api/v1/hello
2. API Accepts POST request with body:
`
{name: 'John Doe'}
`

3. Returns `Hello, John Doe!`
