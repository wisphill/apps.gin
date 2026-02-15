# apps.gin

A production-ready REST API built with Gin demonstrating clean architecture, JWT authentication, structured logging, middleware design, and observability patterns.

## Features

- ✅ RESTful API with Gin
- ✅ JWT Authentication (HS256)
- ✅ Custom JWT Middleware
- ✅ Trace ID propagation
- ✅ Structured logging with Zap
- ✅ Context-aware logging (trace injection)
- ✅ Graceful shutdown
- ✅ Environment config via .env
- ✅ Clean architecture layering
- ✅ Unit testing with httptest
- ✅ Middleware chaining
- ✅ JSON binding & validation

# Project structure

```
apps.gin/
├── cmd/
│   └── main.go
├── internal/
│   ├── logger/
│   ├── middleware/
│   ├── handler/
│   ├── service/
│   ├── repository/
│   └── utils/
├── scripts/
├── go.mod
└── README.md
```

```js
# Generate a token
go run scripts/generate_token.go -user=123 -role=admin
```

### Docker

```
docker build -t apps-gin .
docker run -p 8080:8080 apps-gin
```
