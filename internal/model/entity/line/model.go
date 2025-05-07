package line

import (
	"game_server_slots_fortune_snake/internal/model/entity/symbol"
)

type Item struct {
	Index     int            // 賠付線的位置
	Positions []Position     // 中獎符號位置
	Symbols   []*symbol.Item // 線的組成符號
	Symbol    *symbol.Item   // 中獎符號
	Score     int            // 中獎金額
}

type Position struct {
	Col int
	Row int
}
