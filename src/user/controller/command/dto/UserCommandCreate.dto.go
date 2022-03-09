package userControllerCommandDto

// CreateRequest : Request for CreateUser
type CreateRequest struct {
	NickName string `json:"nickname"`
	Id       string `json:"id"`
	Password string `json:"password"`
}

// CreateResponse : Response for CreateUser
// @Description SUCCESS 만이 성공
// TODO: Insert UserObject
type CreateResponse struct {
	Code string `json:"code" example:"SUCCESS"`
}

// CreateResponseError : Response for CreateUser
// @Description Create 의 오류응답
type CreateResponseError struct {
	Code         string `json:"code"`
	ErrorMessage string `json:"errorMessage"`
}

var CreateResponseErrorForBadRequest = CreateResponseError{
	Code:         "BAD_REQUEST",
	ErrorMessage: "All fields are required",
}

// TODO : Response/Error 공용화하기
