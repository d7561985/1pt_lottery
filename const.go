package lottery

const (
	// After creation cookie live time 2h
	Cookie = "lottery"
)

var (
	WsEventEnter = "enter"
	WsEventLeave = "leave"
	WsList       = "list"
	WsEventBegin = "begin"
	WsEventStop  = "stop"
)
