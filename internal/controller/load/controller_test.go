package load

import (
	weightRepo "game_server_slots_fortune_snake/internal/repository/weight"
	weightService "game_server_slots_fortune_snake/internal/service/weight"
	"testing"
)

func TestLoadController_Load(t *testing.T) {
	repo := weightRepo.New()
	serv := weightService.New(repo)
	conn := New(serv)
	conn.Load()
}
