package response

type UldResponse struct {
	Id           int    `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	AccessToken  string `json:"access_token"`
	CreatedDate  string `json:"created_date"`
	IsActive     bool   `json:"is_Active"`
	RegisteredId int    `json:"registerd_id"`
}
