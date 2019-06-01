package dto

import (
	_ "github.com/d7561985/1pt_lottery/pkg/vld"
)

type UserRequest struct {
	Name   string `json:"user" validate:"required|ascii" filter:"trim"`
	Avatar string `json:"avatar,omitempty" validate:"img2"`
}

type StartResponse struct {
	UUID string `json:"uuid"`
}

type WSEvent struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

type WSNameResponse struct {
	UserRequest
	Competitors int `json:"competitors"`
}

// whisper personal response
type WSListResponse struct {
	Me   string        `json:"me"`
	List []UserRequest `json:"list"`
}
