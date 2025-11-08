package user

type User struct {
	ID string `json:"id"`
	Password string `json:"password"`
	Likes []string
}
