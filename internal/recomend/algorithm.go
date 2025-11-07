package recomend

import "recsys/internal/user"

func Recomend(user user.User, users []user.User) []string {
	liked := make(map[string]bool)
	scores := make(map[string]int)

	for _ ,m := range user.Likes{
		liked[m] = true
	}
	for _, other := range users {
		if other.ID == user.ID {
			continue
		}

		common := 0
		for _ , m := range other.Likes{
			if liked[m] {
				common++
			}
		}
		if common > 0 {
			for _ , m := range other.Likes {
				if !liked[m] {
					scores[m] += common
				}
			}
		}
	}
	var result []string
	for movie := range scores {
		result = append(result, movie)
	}
	return result
}
