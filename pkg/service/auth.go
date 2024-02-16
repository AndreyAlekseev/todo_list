package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	todo "github.com/AndreyAlekseev/todo_list"
	"github.com/AndreyAlekseev/todo_list/pkg/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "earvEAVBAIBNSDHBrvubou"
	signingKey = "evjnweviuniwbD#vibrevjno"
	tokenTTl   = 12 * time.Hour
)

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}
type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	//get user from DB
	user, err := s.repo.GetUser(username, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTl).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id})
	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
