package request

type CreateUserRequest struct {
	Name    string `validate:"required,min=1,max=200" json:"name"`
	Email   string `validate:"required,min=1,max=200" json:"email"`
	Address string `validate:"required,min=1,max=200" json:"address"`
}
