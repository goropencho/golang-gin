package request

type UpdateUserRequest struct {
	Id      int    `validate:"required"`
	Name    string `validate:"required,max=200,min=1" json:"name"`
	Email   string `validate:"required,max=200,min=1" json:"email"`
	Address string `validate:"required,max=200,min=1" json:"address"`
}
