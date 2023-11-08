package middlewares

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

type AuthServiceImpl struct {
}

var SECRET_KEY = []byte("PURPLESNAP_ITSASECRET")

func (a *AuthServiceImpl) GenerateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func (a *AuthServiceImpl) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid token")
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}

func NewAuthService() AuthService {
	return &AuthServiceImpl{}
}
