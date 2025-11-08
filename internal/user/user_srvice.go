package user

import (
	"fmt"
	"math/rand"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

type UserServiec struct {
	repo *UserRepository
}

func NewUserService(UserRepository *UserRepository) *UserServiec {
	return &UserServiec{repo: UserRepository}
}

func (s *UserServiec) RandomUserId() (string, error){
	BaseName := "recsysgo-"
	randomInt := rand.Intn(1000000)
	return BaseName + strconv.Itoa(randomInt), nil
}
func (s *UserServiec) SignNewUser(user *User) (string, error) {
    if user == nil {
        return "", fmt.Errorf("user is nil")
    }

    if user.Password == "" {
        return "", fmt.Errorf("password is required")
    }
    if user.ID == "" {
        id, err := s.RandomUserId()
        if err != nil {
            return "", fmt.Errorf("generate user id: %w", err)
        }
        user.ID = id
    }

    hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return "", fmt.Errorf("hash password: %w", err)
    }
    user.Password = string(hashed)
	if err := s.repo.SaveUser(user); err != nil {
        return "", fmt.Errorf("save new user: %w", err)
    }

    return user.ID, nil
}
