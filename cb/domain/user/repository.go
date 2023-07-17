package user

type UserRepository interface {
	FindAll() (*[]User, error)
	CreateUser(User) (*CreateResponse, error)
	GetByID(string) (*User, error)
	AddCardIdToUser(string, string) error
}
