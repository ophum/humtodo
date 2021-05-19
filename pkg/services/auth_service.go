package services

import (
	"fmt"
	"time"

	"github.com/ophum/humtodo/pkg/entities"
	"github.com/ophum/humtodo/pkg/repositories"
	"golang.org/x/crypto/bcrypt"

	"github.com/dgrijalva/jwt-go"
)

type AuthService struct {
	secret   []byte
	userRepo repositories.UserRepository
}

func NewAuthService(secret []byte, userRepo *repositories.UserRepository) *AuthService {
	return &AuthService{
		secret:   secret,
		userRepo: *userRepo,
	}
}

func (s *AuthService) SignUp(name, password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}

	user, err := s.userRepo.Create(entities.UserEntity{
		Name:     name,
		Password: string(hashed),
	})
	if err != nil {
		return "", err
	}

	return s.generateToken(user.ID, user.Name)
}

func (s *AuthService) SignIn(name, password string) (string, error) {
	user, err := s.userRepo.FindByName(name)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	return s.generateToken(user.ID, name)
}

func (s *AuthService) Verify(token string) (string, error) {
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.secret), nil
	})

	if t.Valid {
		claims, ok := t.Claims.(jwt.MapClaims)
		if !ok {
			return "", fmt.Errorf("unauthorized")
		}

		userId := claims["uid"].(string)
		_, err := s.userRepo.Find(userId)
		if err != nil {
			return "", fmt.Errorf("unauthorized")
		}
		return s.generateToken(claims["uid"].(string), claims["name"].(string))
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		return "", ve
	} else {
		return "", err
	}
}

func (s *AuthService) generateToken(id, name string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = id
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	return token.SignedString(s.secret)
}
