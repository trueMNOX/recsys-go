package movie

import (
	"errors"
	"fmt"
	"recsys/internal/user"
)

type MovieService struct {
	repo     *MovieRepository
	userRepo *user.UserRepository
}

func NewMovieService(repo *MovieRepository, userRepo *user.UserRepository) *MovieService {
	return &MovieService{
		repo:     repo,
		userRepo: userRepo,
	}
}

var ErrMovieAlreadyLiked = errors.New("movie has already been liked")

func (s *MovieService) LikeMovie(userID, movieID string) error {
	movie, err := s.repo.FindMovieByID(movieID)
	if err != nil {
		return fmt.Errorf("finding movie: %w", err)
	}

	user, err := s.userRepo.FindUserById(fmt.Sprint(userID))
	if err != nil {
		return fmt.Errorf("finding user: %w", err)
	}

	for _ , id := range user.Likes{
		if id == movie.ID {
			return ErrMovieAlreadyLiked
		}
	}
	user.Likes = append(user.Likes, movie.ID)
	
	if err := s.userRepo.SaveUser(user); err != nil {
		return fmt.Errorf("updating user likes: %w", err)
	}

	return nil
}
