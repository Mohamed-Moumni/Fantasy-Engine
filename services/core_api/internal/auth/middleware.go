package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

func AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			var ctx context.Context

			fmt.Println("Auth Header: ", authHeader)

			if authHeader != "" {
				tokenString := strings.TrimPrefix(authHeader, "Bearer ")
				if tokenString == "" {
					ctx = context.WithValue(r.Context(), "userID", "")
				}
				userID, err := ValidateJWT(tokenString)
				if err != nil {
					ctx = context.WithValue(r.Context(), "userID", "")
				}
				ctx = context.WithValue(r.Context(), "userID", userID)

			} else {
				ctx = context.WithValue(r.Context(), "userID", "")
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetUserIDFromContext(ctx context.Context) (string, error) {
	userID, ok := ctx.Value("userID").(string)
	if !ok {
		return "", fmt.Errorf("user ID not found in context")
	}
	return userID, nil
}
