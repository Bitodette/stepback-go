package middleware

import (
	"context"
	"net/http"
	"strings"

	"stepback-golang/internal/utils"
)

type contextKey string

const UserIDKey contextKey = "user_id"
const UserRoleKey contextKey = "user_role"

// Auth validates JWT from Authorization header,
// extracts user_id and role into request context
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			utils.Error(w, http.StatusUnauthorized, "Missing authorization header")
			return
		}

		// format: "Bearer <token>"
		parts := strings.SplitN(header, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.Error(w, http.StatusUnauthorized, "Invalid authorization format")
			return
		}

		claims, err := utils.ValidateAccessToken(parts[1])
		if err != nil {
			utils.Error(w, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		// inject user info ke context, biar handler bisa akses
		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
		ctx = context.WithValue(ctx, UserRoleKey, claims.Role)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserID gets user ID from context
// NOTE: only works if Auth middleware ran first
func GetUserID(r *http.Request) uint {
	id, _ := r.Context().Value(UserIDKey).(uint)
	return id
}

func GetUserRole(r *http.Request) string {
	role, _ := r.Context().Value(UserRoleKey).(string)
	return role
}
