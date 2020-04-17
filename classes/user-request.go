package classes

type UserRequestPost struct {
	Email		string		`json:"email"`
	Password	string		`json:"password"`
}

type UserRequestPut struct {
	UserRequestPost
}