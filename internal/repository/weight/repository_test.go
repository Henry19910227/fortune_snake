package weight

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository_LoadBaseWeight(t *testing.T) {
	repo := New()
	repo.LoadBaseWeight()
	for _, stat := range repo.BaseWeight() {
		fmt.Println(stat)
	}
	assert.Equal(t, 148, len(repo.BaseWeight()))
}

func TestRepository_LoadFreeWeight(t *testing.T) {
	repo := New()
	repo.LoadFreeWeight()
	for _, stat := range repo.FreeWeight() {
		fmt.Println(stat)
	}
	assert.Equal(t, 224, len(repo.FreeWeight()))
}
