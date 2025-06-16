package bet

type Response struct {
	GameMode   string      `json:"game_mode"`
	Balance    int         `json:"balance"` // 真實餘額
	GameResult *GameResult `json:"game_result,omitempty"`
}

func NewResponse(gameMode string) *Response {
	response := &Response{}
	response.GameMode = gameMode
	return response
}
