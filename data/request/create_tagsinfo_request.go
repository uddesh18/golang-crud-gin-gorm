package request

type CreateTagsInfoRequest struct {
	Tid    int    `validate:"required,min=1,max=200" json:"name"`
	Class  string `validate:"required,min=1,max=200" json:"lname"`
	Bg     string `validate:"required,min=1,max=200" json:"bg"`
	Dob    string `validate:"required,min=1,max=200" json:"dob"`
	Gender int    `validate:"required,min=1,max=200" json:"gender"`
}
