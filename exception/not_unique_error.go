package exception

type NotUniqueError struct {
	Error string
}

func NewNotUniqueError() NotUniqueError {
	return NotUniqueError{Error: "User Email Already Exist"}
}
