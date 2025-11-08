package movie

import (
	"fmt"

	"gorm.io/gorm"
)

type MovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepository{
	return &MovieRepository{db: db}
}

func (r *MovieRepository) FindMovieByID(MovieID string) (*Movie, error){
	var movie Movie
	err := r.db.Where("ID = ? ", MovieID).First(&movie).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Errorf("Movie with this ID: %s Not found in Database", MovieID)
		}
		return nil, err
	}
	return &movie, nil
}
