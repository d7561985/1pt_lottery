package dto

type StartRequest struct {
	User string `json:"user" validate:"required|ascii" filter:"trim"`
}

type StartResponse struct {
	UUID string `json:"uuid"`
}

type WSEvent struct {
	Event string `json:"event"`
}

type WSNameResponse struct {
	WSEvent
	Name        string `json:"name"`
	Competitors int    `json:"competitors"`
}

type WSListResponse struct {
	WSEvent
	List []string `json:"list"`
}
