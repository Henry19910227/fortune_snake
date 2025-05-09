package reels_free

import (
	"game_server_slots_fortune_snake/internal/model/entity/symbol"
	symbolRepo "game_server_slots_fortune_snake/internal/repository/symbol"
	"math/rand"
)

type service struct {
	symbolRepo symbolRepo.Repository
}

func New(symbolRepo symbolRepo.Repository) Service {
	return &service{symbolRepo: symbolRepo}
}

func (s *service) Generate(layout []int) [][][]*symbol.Item {
	// 選擇一個隨機符號
	randomSymbol := s.symbolRepo.GetRandomSymbol()
	reelsList := make([][][]*symbol.Item, 0)
	reelsList = append(reelsList, s.generate(randomSymbol, layout))
	s.handle(randomSymbol, &reelsList, layout)
	return reelsList
}

func (s *service) handle(randomSymbol *symbol.Item, reelsList *[][][]*symbol.Item, layout []int) {
	lastReels := (*reelsList)[len(*reelsList)-1]
	targetReels := make([][]*symbol.Item, len(lastReels))
	s.deepCopy(targetReels, lastReels)
	coverReels := s.generate(randomSymbol, layout)
	s.cover(targetReels, coverReels)
	if s.isEqual(lastReels, targetReels) {
		*reelsList = append(*reelsList, targetReels)
		return
	}
	*reelsList = append(*reelsList, targetReels)
	s.handle(randomSymbol, reelsList, layout)
}

func (s *service) generate(randomSymbol *symbol.Item, layout []int) [][]*symbol.Item {
	// 獲取一個空白符號
	spaceSymbol := s.symbolRepo.GetSymbol(99)
	// ㄧ ~ 三軸
	reels := make([][]*symbol.Item, 0)
	for col := 0; col < len(layout); col++ {
		reel := make([]*symbol.Item, 0)
		for row := 0; row < layout[col]; row++ {
			num := rand.Intn(spaceSymbol.Weight + randomSymbol.Weight)
			if num < randomSymbol.Weight {
				reel = append(reel, randomSymbol)
				continue
			}
			reel = append(reel, spaceSymbol)
		}
		reels = append(reels, reel)
	}
	// 將第二軸覆蓋為全 wild + 空格符號，空格與 wild 符號出現機率各 50 %
	for i := 0; i < len(reels[1]); i++ {
		num := rand.Intn(100) + 1
		if num > 50 {
			reels[1][i] = spaceSymbol
			continue
		}
		reels[1][i] = s.symbolRepo.GetSymbol(0)
	}
	return reels
}

func (s *service) isEqual(first [][]*symbol.Item, second [][]*symbol.Item) bool {
	for i := 0; i < len(first); i++ {
		for j := 0; j < len(first[i]); j++ {
			if first[i][j].ID == second[i][j].ID {
				continue
			}
			return false
		}
	}
	return true
}

func (s *service) cover(dst [][]*symbol.Item, src [][]*symbol.Item) {
	for col := 0; col < len(dst); col++ {
		for row := 0; row < len(dst[col]); row++ {
			if dst[col][row].ID != 99 {
				continue
			}
			dst[col][row].ID = src[col][row].ID
		}
	}
}

func (s *service) deepCopy(dst [][]*symbol.Item, src [][]*symbol.Item) {
	for i := range src {
		dst[i] = make([]*symbol.Item, len(src[i])) // 建立內層
		for j := range src[i] {
			if src[i][j] != nil {
				itemCopy := *src[i][j] // 複製值
				dst[i][j] = &itemCopy  // 指向新值
			}
		}
	}
}
