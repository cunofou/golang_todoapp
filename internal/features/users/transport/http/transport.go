package users_transport_http

type UsersHTTPHandler struct {
	UsersService UsersService
}

type UsersService interface {
}

func NewUsersHTTPHandler(usersService UsersService) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		UsersService: usersService,
	}
}
