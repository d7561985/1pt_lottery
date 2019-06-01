package lottery

const (
	// After creation cookie live time 2h
	Cookie = "lottery"
)

var (
	EventError        = "error"
	WsEventEnter      = "enter"
	WsEventLeave      = "leave"
	WsList            = "list"
	WsEventBegin      = "begin"
	WsEventStop       = "winner"
	WSEventServerTime = "time"
)
