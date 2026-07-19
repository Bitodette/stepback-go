# Stepback E-commerce

rewrite of stepback e-commerce backend in Go, primarily to learn Go properly.

the old backend was built with Django, this time im rebuilding it with Go + PostgreSQL with better architecture and cleaner code.

**tech stack:**
- Go + net/http (chi router)
- PostgreSQL + GORM
- JWT auth (access + refresh token)
- Docker support

**status: in progress**
- [x] config, database, models
- [x] auth (register, login, refresh token)
- [ ] user profile & address management
- [ ] products & categories
- [ ] shopping cart
- [ ] orders & checkout
- [ ] payment (midtrans)
- [ ] shipping (biteship)
- [ ] reviews
- [ ] admin endpoints
- [ ] tests

hopefully i finish this one lol

## running locally

```bash
# setup
cp .env.example .env
# edit .env with your postgres credentials

# run
go run ./cmd/server
```

server starts at `http://localhost:8080`

## api endpoints

| method | endpoint | description |
|--------|----------|-------------|
| POST | /api/users/register | register new user |
| POST | /api/users/login | login, get tokens |
| POST | /api/users/token/refresh | refresh access token |
| GET | /api/users/profile | get profile (auth required) |
| GET | /api/ping | health check |

## bruno collection

open the `bruno/` folder in Bruno desktop to test the API locally.
