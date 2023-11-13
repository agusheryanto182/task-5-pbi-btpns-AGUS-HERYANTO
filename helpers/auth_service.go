package helpers

import "github.com/golang-jwt/jwt/v5"

type AuthService interface {
	GenerateToken(userID int) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}
