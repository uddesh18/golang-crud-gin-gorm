package request

type CreateTagsInfoRequest struct {
	Tid    int    "validate:\"required,min=1,max=200\" json:\"tid\""
	Class  string `validate:"required,min=1,max=200" json:"name"`
	Bg     string `validate:"required,min=1,max=200" json:"bg"`
	Dob    string `validate:"required,min=1,max=200" json:"dob"`
	Gender int    `validate:"required,min=1,max=200" json:"gender"`
}
