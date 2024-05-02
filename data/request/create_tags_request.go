package request

type CreateTagsRequest struct {
	Name     string `validate:"required,min=1,max=200" json:"name"`
	LastName string `validate:"required,min=1,max=200" json:"lname"`
}
