package bet

type Response struct {
	GameMode   string
	Balance    int // 真實餘額
	GameResult *GameResult
}

func NewResponse(gameMode string) *Response {
	response := &Response{}
	response.GameMode = gameMode
	return response
}
