package slot

import (
	"fmt"
	symbolRepo "game_server_slots_fortune_snake/internal/repository/symbol"
	"testing"
)

func TestService_Generate(*testing.T) {
	repo := symbolRepo.New()
	svc := NewService(repo)
	reelSet := svc.Generate([]int{3, 4, 3})
	for row := 0; row < len(reelSet); row++ {
		for col := 0; col < len(reelSet[row]); col++ {
			fmt.Println(reelSet[row][col])
		}
		fmt.Println("-------")
	}
}
