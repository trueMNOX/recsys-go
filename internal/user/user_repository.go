package user

import (
	"fmt"

	"gorm.io/gorm"
)

type userRepository struct{
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db : db}
}

func (r *userRepository) FindUserById(userid string) (*User, error){
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

func (r *userRepository) LikeMovie (userID ,movieID string) (string, error){
	user , err := r.FindUserById(userID)
	if err != nil {
		return "Error", nil
	}
	
}
