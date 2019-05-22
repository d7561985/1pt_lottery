package dto

type StartRequest struct {
	User string `json:"user"`
}

type StartResponse struct {
	UUID string `json:"uuid"`
}
