package middleware

import (
	"fmt"
	"strings"
	"todo-cognixus/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func ValidateToken() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Retrieve token from authorization header
		authHeader := ctx.Get("Authorization")
		if authHeader == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "Missing authorization header")
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid authorization header")
		}

		// Retrieve token from JWT
		token, err := jwt.Parse(parts[1], JWTKeyFunc)
		if err != nil || !token.Valid {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid JWT token")
		}

		// Retrieve claims from token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to extract claims from token")
		}

		// Retrieve user id from claims
		userID, ok := claims["user_id"]
		if !ok {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to extract user ID from claims")
		}

		ctx.Locals("user_id", userID)

		return ctx.Next()
	}
}

func JWTKeyFunc(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return []byte(config.JWT_SECRET_KEY), nil
}
