package model

type Errors struct{}

func (e Error) Error() string {
	return e.String()
}
