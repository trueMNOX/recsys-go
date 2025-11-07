package user

type UserHandler struct {
	service *UserServiec
}

func NewUserHanler(userService *UserServiec) *UserHandler{
	return &UserHandler{service: userService}
}
