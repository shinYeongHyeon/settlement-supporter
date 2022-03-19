package groupControllerCommandDto

// CreateRequest : Request for CreateGroup
type CreateRequest struct {
	Title string `json:"title"`
}

// CreateResponse : Response for CreateGroup
// @Description SUCCESS 만이 성공
// TODO: Insert groupObject
type CreateResponse struct {
	Code string `json:"code" example:"SUCCESS"`
}

// CreateResponseError : Response for CreateGroup
// @Description Create 의 오류응답
type CreateResponseError struct {
	Code         string `json:"code"`
	ErrorMessage string `json:"errorMessage"`
}

var CreateResponseErrorForBadRequest = CreateResponseError{
	Code:         "BAD_REQUEST",
	ErrorMessage: "All fields are required",
}
