package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT creates a JWT token for a given email.
func GenerateJWT(email string) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET"))
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(72 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// GenerateResetToken creates a password reset token (expires in 1 hour).
func GenerateResetToken(email string) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET"))
	claims := jwt.MapClaims{
		"email": email,
		"reset": true,
		"exp":   time.Now().Add(1 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// ValidateJWT parses and validates a JWT token, returning the email claim if valid.
func ValidateJWT(tokenStr string) (string, error) {
	secretKey := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return secretKey, nil
	})
	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		email, ok := claims["email"].(string)
		if !ok {
			return "", fmt.Errorf("invalid token claims")
		}
		return email, nil
	}

	return "", fmt.Errorf("invalid token")
}
