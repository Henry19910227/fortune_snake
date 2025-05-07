package symbol

import (
	"fmt"
	"testing"
)

func TestRepository_GetSymbols(*testing.T) {
	repo := New()
	for _, symbol := range repo.GetSymbols() {
		fmt.Println(symbol)
	}
	fmt.Println(repo.GetTotalWeight())
}
