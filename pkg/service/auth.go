package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/VladislavEF/todo-app"
	"github.com/VladislavEF/todo-app/pkg/repository"
)

const salt = "dgergewtg3453wg3e46y34gh"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generateUPasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generateUPasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
