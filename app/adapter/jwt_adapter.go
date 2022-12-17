package adapter

import (
	"github.com/golang-jwt/jwt/v4"
	"gokomodo-assesment/app/domain/repository"
	"gokomodo-assesment/internal/config"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type JWTAdapter struct {
	SecretKey []byte
}

func NewJWTAdapter() repository.JWTRepository {
	return &JWTAdapter{
		SecretKey: []byte(config.GetConfig().EnvSystem.SecretKey.Value),
	}
}

func (a *JWTAdapter) CheckPassword(userPassword, loginPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(loginPassword))
	if err != nil {
		return err
	}

	return nil
}

func (a *JWTAdapter) GenerateJWT(email, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  email,
		"role": role,
		"exp":  time.Now().Add(time.Hour * 3).Unix(),
	})

	tokenString, err := token.SignedString(a.SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
