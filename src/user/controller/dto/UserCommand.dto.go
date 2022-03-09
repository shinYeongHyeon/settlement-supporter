package userControllerDto

type CreateRequest struct {
	NickName string `json:"nickname"`
	Id       string `json:"id"`
	Password string `json:"password"`
}
