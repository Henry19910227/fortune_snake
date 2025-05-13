package load

import weightService "game_server_slots_fortune_snake/internal/service/weight"

type controller struct {
	weightService weightService.Service
}

func New(weightService weightService.Service) Controller {
	return &controller{weightService: weightService}
}

func (c *controller) Load() {
	c.weightService.Load()
}
