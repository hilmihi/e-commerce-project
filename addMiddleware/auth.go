package addmiddleware

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTService interface {
	GenerateToken(userId int) (string, error)
	ValidateToken(TokenGenerated string) (*jwt.Token, error)
}

type jwtService struct{}

var JWT_SECRET = []byte("BUBURDEPOK")

func AuthService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(JWT_SECRET)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func (s *jwtService) ValidateToken(TokenGenerated string) (*jwt.Token, error) {
	token1, err := jwt.Parse(TokenGenerated, func(token1 *jwt.Token) (interface{}, error) {
		_, ok := token1.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(JWT_SECRET), nil
	})
	if err != nil {
		return token1, err
	}
	return token1, nil
}
