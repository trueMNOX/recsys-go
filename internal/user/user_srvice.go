package user

import()

type UserServiec struct {
	repo *userRepository
}

func NewUserService(userRepository *userRepository) *UserServiec {
	return &UserServiec{repo: userRepository}
}

func (s *UserServiec) Register(userID string) (*User, error){
	
}
