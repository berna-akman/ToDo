package user

type UserRepository interface {
	FindAll() (*[]User, error)
	CreateUser(User) (*CreateResponse, error)
}
