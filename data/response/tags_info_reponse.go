package response

type TagsInfoResponse struct {
	Tid    int    `json:"id"`
	Class  string `json:"class"`
	Bg     string `json:"bloodGroup"`
	Dob    string `json:"dob"`
	Gender int    `json:"gender"`
}
