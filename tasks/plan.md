# Spec: Stepback E-commerce Backend Rewrite

## Objective

Rewrite backend e-commerce Stepback dari Django (Python) ke Go dengan `net/http` standard library. Fase pertama: Authentication & User Profile. Database: PostgreSQL dengan GORM.

## Tech Stack

- **Language:** Go 1.22+
- **HTTP Router:** `net/http` + `chi` (lightweight router, stdlib-compatible)
- **Database:** PostgreSQL 16
- **ORM:** GORM v2
- **Auth:** JWT (golang-jwt/jwt/v5), bcrypt for password hashing
- **Config:** godotenv + env var
- **Migration:** GORM AutoMigrate (development), golang-migrate (production)

## Commands

```bash
# Build
go build -o bin/server ./cmd/server

# Run
go run ./cmd/server

# Test
go test ./... -v -race

# Lint
golangci-lint run ./...

# Dev (with air for hot reload)
air -c .air.toml
```

## Project Structure

```
stepback-golang/
├── cmd/
│   └── server/
│       └── main.go              # Entry point
├── internal/
│   ├── config/
│   │   └── config.go            # Env var loading
│   ├── database/
│   │   ├── db.go                # GORM connection
│   │   └── migrate.go           # AutoMigrate
│   ├── middleware/
│   │   ├── auth.go              # JWT auth middleware
│   │   ├── cors.go              # CORS middleware
│   │   └── logger.go            # Request logging
│   ├── model/
│   │   ├── user.go              # User model
│   │   └── address.go           # Address model
│   ├── handler/
│   │   ├── auth.go              # Register, Login, Logout, Refresh, Verify Email, Password Reset
│   │   ├── user.go              # Profile, Update Profile, Change Password
│   │   └── address.go           # CRUD Address
│   ├── repository/
│   │   ├── user.go              # User DB operations
│   │   └── address.go           # Address DB operations
│   ├── service/
│   │   ├── auth.go              # Auth business logic
│   │   └── user.go              # User business logic
│   ├── router/
│   │   └── router.go            # Route definitions
│   └── utils/
│       ├── response.go          # Standard JSON response helpers
│       ├── jwt.go               # JWT token generation/validation
│       └── validation.go        # Input validation helpers
├── migrations/
│   └── 001_initial.up.sql       # SQL migrations (for production)
├── tasks/
│   ├── plan.md                  # This spec
│   └── todo.md                  # Task checklist
├── .env.example
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## Code Style

```go
// Handler example - clean, consistent error handling
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
    var req RegisterRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        utils.Error(w, http.StatusBadRequest, "Invalid request body")
        return
    }

    if err := h.validator.Struct(req); err != nil {
        utils.ValidationErrors(w, err)
        return
    }

    user, err := h.service.Register(r.Context(), &req)
    if err != nil {
        utils.Error(w, http.StatusConflict, err.Error())
        return
    }

    utils.Created(w, user)
}
```

**Conventions:**
- Exported functions: PascalCase (Go standard)
- Unexported: camelCase
- Files: snake_case (Go convention for non-test files)
- Handler methods: verb + noun (e.g., `Register`, `Login`, `GetProfile`)
- Repository methods: CRUD pattern (e.g., `FindByEmail`, `Create`)
- Response format: always `{ "success": bool, "message": string, "data": any }`
- Errors: always return JSON, never plain text

## Response Format

```json
// Success
{
  "success": true,
  "message": "Operation successful",
  "data": { }
}

// Error
{
  "success": false,
  "message": "Error description",
  "errors": { "field": ["error detail"] }
}

// Pagination
{
  "success": true,
  "count": 100,
  "next": "http://api.example.com/path?limit=12&offset=12",
  "previous": null,
  "data": []
}
```

## Testing Strategy

- **Framework:** Go standard `testing` + `httptest`
- **Unit tests:** Service layer (business logic)
- **Integration tests:** Repository layer (against real/test DB)
- **Handler tests:** HTTP request/response with `httptest`
- **Coverage target:** >70% for new code
- **Test location:** `*_test.go` next to source files

## Phase 1 Scope (Auth + User)

### Database Schema

```sql
-- Users
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20),
    role VARCHAR(20) DEFAULT 'customer',  -- customer, admin, staff
    is_verified BOOLEAN DEFAULT FALSE,
    verification_token VARCHAR(255),
    reset_token VARCHAR(255),
    reset_token_expires_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Addresses
CREATE TABLE addresses (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    label VARCHAR(50) NOT NULL,         -- Home, Office, etc.
    recipient_name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    address_line TEXT NOT NULL,
    city VARCHAR(100) NOT NULL,
    province VARCHAR(100) NOT NULL,
    postal_code VARCHAR(10) NOT NULL,
    is_default BOOLEAN DEFAULT FALSE,
    biteship_area_id VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
```

### API Endpoints (Phase 1)

```
POST   /api/users/register
POST   /api/users/login
POST   /api/users/logout
POST   /api/users/token/refresh
POST   /api/users/verify-email
POST   /api/users/request-password-reset
POST   /api/users/reset-password-confirm

GET    /api/users/profile
PUT    /api/users/profile/update
POST   /api/users/change-password

GET    /api/users/addresses
POST   /api/users/addresses
PUT    /api/users/addresses/{id}
DELETE /api/users/addresses/{id}
```

## Boundaries

- **Always:** Run tests before commits, validate inputs, return JSON errors, use context for request-scoped values
- **Ask first:** Database schema changes, adding new dependencies, changing JWT secret location
- **Never:** Commit secrets (.env), skip auth middleware on protected routes, return raw errors to client

## Success Criteria

- [ ] All Phase 1 endpoints functional and tested
- [ ] JWT auth working (access + refresh tokens)
- [ ] Password hashing with bcrypt
- [ ] Email verification flow (token-based)
- [ ] Password reset flow (token-based)
- [ ] Address CRUD with user isolation
- [ ] Consistent JSON response format
- [ ] Graceful error handling (no panics, no raw DB errors)

## Open Questions

1. Untuk email verification & password reset, mau pakai email service apa? (SendGrid, SMTP langsung, atau mock dulu?)
2. JWT access token expiry berapa lama? (default: 15 menit)
3. Refresh token expiry berapa lama? (default: 7 hari)
