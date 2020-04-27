package classes

type UserRequestCreate struct {
	Email		string		`form:"email" json:"email" validate:"required,email"`
	Password	string		`form:"password" json:"password" validate:"required,min=8,max=32"`
}

type UserRequestUpdate struct {
	Email		string		`form:"email" json:"email" validate:"required,email"`
	Password	string		`form:"password" json:"password" validate:"required,min=8,max=32"`
}