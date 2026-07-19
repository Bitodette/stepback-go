# Todo: Phase 1 - Auth + User Profile

## Phase 1.1: Project Setup
- [ ] Initialize Go module (`go mod init stepback-golang`)
- [ ] Install dependencies (chi, gorm, jwt, bcrypt, godotenv)
- [ ] Create project structure (cmd/, internal/, migrations/)
- [ ] Setup `.env.example` and `config/config.go`
- [ ] Setup `database/db.go` (PostgreSQL connection via GORM)
- [ ] Setup `cmd/server/main.go` (entry point)
- [ ] Setup `.gitignore`
- [ ] Verify: server starts and connects to DB

## Phase 1.2: Models & Migration
- [ ] Create `model/user.go` (User struct)
- [ ] Create `model/address.go` (Address struct)
- [ ] Setup `database/migrate.go` (AutoMigrate)
- [ ] Verify: tables created in PostgreSQL

## Phase 1.3: Utils & Middleware
- [ ] Create `utils/response.go` (JSON response helpers)
- [ ] Create `utils/jwt.go` (token generation/validation)
- [ ] Create `utils/validation.go` (input validation)
- [ ] Create `middleware/auth.go` (JWT extraction & context injection)
- [ ] Create `middleware/cors.go`
- [ ] Create `middleware/logger.go`
- [ ] Verify: middleware chain works

## Phase 1.4: Auth - Register & Login
- [ ] Create `repository/user.go` (FindByEmail, Create)
- [ ] Create `service/auth.go` (Register, Login business logic)
- [ ] Create `handler/auth.go` (Register, Login handlers)
- [ ] Setup `router/router.go` (POST /register, POST /login)
- [ ] Verify: can register user and login, get JWT tokens

## Phase 1.5: Auth - Token Refresh & Logout
- [ ] Implement POST /token/refresh endpoint
- [ ] Implement POST /logout endpoint (token blacklist or just invalidate)
- [ ] Verify: refresh token works, logout invalidates

## Phase 1.6: Auth - Email Verification & Password Reset
- [ ] Implement POST /verify-email endpoint
- [ ] Implement POST /request-password-reset endpoint
- [ ] Implement POST /reset-password-confirm endpoint
- [ ] Decide: mock email or real service (start with mock/log)
- [ ] Verify: full verification and reset flow

## Phase 1.7: User Profile
- [ ] Create `service/user.go` (GetProfile, UpdateProfile, ChangePassword)
- [ ] Create `handler/user.go` (GET /profile, PUT /profile/update, POST /change-password)
- [ ] Add routes to router
- [ ] Verify: profile CRUD works with auth

## Phase 1.8: Address Management
- [ ] Create `repository/address.go` (CRUD operations)
- [ ] Create `handler/address.go` (CRUD handlers)
- [ ] Add routes to router
- [ ] Verify: address CRUD with user isolation

## Phase 1.9: Testing & Cleanup
- [ ] Write unit tests for service layer
- [ ] Write handler tests with httptest
- [ ] Run golangci-lint, fix issues
- [ ] Update README.md
- [ ] Verify: all tests pass, no lint errors
