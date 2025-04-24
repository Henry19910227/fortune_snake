package model

type MessageRequest struct {
	Action   string
	PlayerId uint64
	Data     string
}

type MessageResponse struct {
	Code    int32
	Message string
	Data    string
}
