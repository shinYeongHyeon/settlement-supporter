package userControllerDto

// CreateRequest : Request for CreateUser
type CreateRequest struct {
	NickName string `json:"nickname"`
	Id       string `json:"id"`
	Password string `json:"password"`
}

// CreateResponse : Response for CreateUser
type CreateResponse struct {
	Code string `json:"code"`
}
