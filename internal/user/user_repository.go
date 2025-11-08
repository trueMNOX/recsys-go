package user

import (
	"fmt"

	"gorm.io/gorm"
)

type UserRepository struct{
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db : db}
}

func (r *UserRepository) FindUserById(userid string) (*User, error){
	var user User
	err := r.db.Where("ID = ?", userid).First(&user).Error;
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Errorf("User with this ID: %s not found in Database", userid)
		}
		return nil , err
	}
	return &user, nil
}

func (r *UserRepository) SaveUser(user *User) error{
	if user == nil {
		return fmt.Errorf("user is empty")
	}
	if err := r.db.Save(user).Error; err != nil {
		return fmt.Errorf("save user: %w", err)
	}
	return nil 
}
