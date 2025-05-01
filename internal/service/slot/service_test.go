package slot

import (
	"fmt"
	symbolRepo "game_server_slots_fortune_snake/internal/repository/symbol"
	"testing"
)

func TestService_Generate(*testing.T) {
	repo := symbolRepo.New()
	svc := NewService(repo)
	fmt.Println(svc.Generate([]int{3, 4, 3}))
}
