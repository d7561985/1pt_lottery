package dto

type StartRequest struct {
	User string `json:"user" validate:"required|ascii" filter:"trim"`
}

type StartResponse struct {
	UUID string `json:"uuid"`
}

type WSEvent struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

type WSNameResponse struct {
	Name        string `json:"name"`
	Competitors int    `json:"competitors"`
}

// whisper personal response
type WSListResponse struct {
	Me   string   `json:"me"`
	List []string `json:"list"`
}
