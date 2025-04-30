package weight

import (
	"fmt"
	"testing"
)

func TestService_LoadBaseWeight(*testing.T) {
	repo := New()
	repo.LoadBaseWeight()
	for _, stat := range repo.BaseWeight() {
		fmt.Println(stat)
	}
}
