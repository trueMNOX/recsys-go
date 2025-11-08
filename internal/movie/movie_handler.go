package movie

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MovieHandler struct {
	Service *MovieService
}

func NewMovieHandler(Service *MovieService)*MovieHandler{
	return &MovieHandler{Service: Service}
}
func (h *MovieHandler) MovieRoute(r *gin.RouterGroup){
	r.GET("/users/:userID/movie/:movieID/like", h.LikeMovie)
}

func (h *MovieHandler) LikeMovie(c *gin.Context) {
    userID := c.Param("userID")
    movieID := c.Param("movieID")

    if userID == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "missing user ID",
        })
        return
    }

    if movieID == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "missing movie ID",
        })
        return
    }

    err := h.Service.LikeMovie(userID, movieID)
    if err != nil {
        switch {
        case errors.Is(err, ErrMovieAlreadyLiked):
            c.JSON(http.StatusConflict, gin.H{
                "error": "movie already liked",
            })
        case errors.Is(err, gorm.ErrRecordNotFound):
            c.JSON(http.StatusNotFound, gin.H{
                "error": "movie or user not found",
            })
        default:
            c.JSON(http.StatusInternalServerError, gin.H{
                "error": "failed to like movie",
            })
        }
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status": "movie liked successfully",
        "userID": userID,
        "movieID": movieID,
    })
}
