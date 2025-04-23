package player

import "game_server_slots_fortune_snake/server"

type controller struct {
}

func New() Controller {
	return &controller{}
}

func (c *controller) GetPlayerSession(ctx *server.Context) {
	//TODO implement me
	panic("implement me")
}
