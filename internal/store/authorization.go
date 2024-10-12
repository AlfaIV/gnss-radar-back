package store

type IAuthorizationStore interface {
	Signin(username, password string) (string, error)
	Signup(username, password string) (string, error)
	Authcheck(value string) (bool, error)
	Logout(value string) (bool, error)
}

type AuthorizationStore struct {
	storage *Storage
}

func (a AuthorizationStore) Signin(username, password string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorizationStore) Signup(username, password string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorizationStore) Authcheck(value string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthorizationStore) Logout(value string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewAuthorizationStore(storage *Storage) *AuthorizationStore {
	return &AuthorizationStore{
		storage: storage,
	}
}
