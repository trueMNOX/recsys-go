package user

import()

type UserServiec struct {
	repo *UserRepository
}

func NewUserService(UserRepository *UserRepository) *UserServiec {
	return &UserServiec{repo: UserRepository}
}
