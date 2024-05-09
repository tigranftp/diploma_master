package jwt_token

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	tokenTTL        = 5 * time.Minute
	refreshTokenTTL = 5 * 24 * time.Hour
)

type refreshTokenClaims struct {
	jwt.StandardClaims
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
}

func ParseRefreshToken(refreshToken string) error {
	_, err := jwt.ParseWithClaims(refreshToken, &refreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingRefreshKey), nil
	})
	return err
}

func ParseToken(accessToken string) (int64, string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, "", err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, "", errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, claims.Username, nil
}

func GenerateTokens(userID int64, username string) (string, string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userID,
		username,
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &refreshTokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(refreshTokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	ttk, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", "", err
	}
	rttk, err := refreshToken.SignedString([]byte(signingRefreshKey))
	return ttk, rttk, err
}
