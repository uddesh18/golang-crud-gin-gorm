package request

type UserRegistrationRequest struct {
	Username string `validate:"required,min=1,max=200" json:"username"`
	Password string `validate:"required,min=1,max=200" json:"password"`
	Email    string `validtea:"required,min=1,max=200" json:"email"`
	Role     string `validtea:"required,min=1,max=200" json:"role"`
}
