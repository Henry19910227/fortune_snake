package bet

type Response struct {
	GameMode   string
	Gold       int // 真實餘額
	ScoreTry   int // 試玩餘額
	GameResult *GameResult
}

func NewResponse(gameMode string) *Response {
	response := &Response{}
	response.GameMode = gameMode
	return response
}
