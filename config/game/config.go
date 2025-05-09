package game

import model "game_server_slots_fortune_snake/internal/model/config/bucket"

type config struct {
}

func New() Config {
	return &config{}
}

func (c *config) BaseBucketConfig() []*model.Item {
	return []*model.Item{
		{LowerLimit: 0, UpperLimit: 0, MaxCapacity: 5000},
		{LowerLimit: 0, UpperLimit: 0.2, MaxCapacity: 0},
		{LowerLimit: 0.2, UpperLimit: 0.3, MaxCapacity: 200},
		{LowerLimit: 0.3, UpperLimit: 0.4, MaxCapacity: 0},
		{LowerLimit: 0.4, UpperLimit: 0.5, MaxCapacity: 200},
		{LowerLimit: 0.5, UpperLimit: 0.6, MaxCapacity: 200},
		{LowerLimit: 0.6, UpperLimit: 0.7, MaxCapacity: 0},
		{LowerLimit: 0.7, UpperLimit: 0.8, MaxCapacity: 200},
		{LowerLimit: 0.8, UpperLimit: 0.9, MaxCapacity: 30},
		{LowerLimit: 0.9, UpperLimit: 1, MaxCapacity: 200},
		{LowerLimit: 1, UpperLimit: 2, MaxCapacity: 200},
		{LowerLimit: 2, UpperLimit: 3, MaxCapacity: 200},
		{LowerLimit: 3, UpperLimit: 4, MaxCapacity: 200},
		{LowerLimit: 4, UpperLimit: 5, MaxCapacity: 200},
		{LowerLimit: 5, UpperLimit: 6, MaxCapacity: 200},
		{LowerLimit: 6, UpperLimit: 7, MaxCapacity: 200},
		{LowerLimit: 7, UpperLimit: 8, MaxCapacity: 200},
		{LowerLimit: 8, UpperLimit: 9, MaxCapacity: 200},
		{LowerLimit: 9, UpperLimit: 10, MaxCapacity: 200},
		{LowerLimit: 10, UpperLimit: 15, MaxCapacity: 150},
		{LowerLimit: 15, UpperLimit: 20, MaxCapacity: 120},
		{LowerLimit: 20, UpperLimit: 30, MaxCapacity: 100},
		{LowerLimit: 30, UpperLimit: 40, MaxCapacity: 90},
		{LowerLimit: 40, UpperLimit: 50, MaxCapacity: 20},
		{LowerLimit: 50, UpperLimit: 100, MaxCapacity: 10},
		{LowerLimit: 100, UpperLimit: 200, MaxCapacity: 10},
		{LowerLimit: 200, UpperLimit: 500, MaxCapacity: 5},
		{LowerLimit: 500, UpperLimit: 1000, MaxCapacity: 2},
		{LowerLimit: 1000, UpperLimit: 5000, MaxCapacity: 1},
	}
}

func (c *config) FreeBucketConfig() []*model.Item {
	return []*model.Item{
		{LowerLimit: 0, UpperLimit: 2, MaxCapacity: 30},
		{LowerLimit: 2, UpperLimit: 4, MaxCapacity: 30},
		{LowerLimit: 4, UpperLimit: 6, MaxCapacity: 30},
		{LowerLimit: 6, UpperLimit: 8, MaxCapacity: 30},
		{LowerLimit: 8, UpperLimit: 10, MaxCapacity: 30},
		{LowerLimit: 10, UpperLimit: 15, MaxCapacity: 100},
		{LowerLimit: 15, UpperLimit: 20, MaxCapacity: 100},
		{LowerLimit: 20, UpperLimit: 25, MaxCapacity: 100},
		{LowerLimit: 25, UpperLimit: 30, MaxCapacity: 100},
		{LowerLimit: 30, UpperLimit: 35, MaxCapacity: 50},
		{LowerLimit: 35, UpperLimit: 40, MaxCapacity: 50},
		{LowerLimit: 40, UpperLimit: 45, MaxCapacity: 50},
		{LowerLimit: 45, UpperLimit: 50, MaxCapacity: 30},
		{LowerLimit: 50, UpperLimit: 55, MaxCapacity: 0},
		{LowerLimit: 55, UpperLimit: 60, MaxCapacity: 30},
		{LowerLimit: 60, UpperLimit: 65, MaxCapacity: 0},
		{LowerLimit: 65, UpperLimit: 70, MaxCapacity: 30},
		{LowerLimit: 70, UpperLimit: 75, MaxCapacity: 0},
		{LowerLimit: 75, UpperLimit: 80, MaxCapacity: 30},
		{LowerLimit: 80, UpperLimit: 85, MaxCapacity: 0},
		{LowerLimit: 85, UpperLimit: 90, MaxCapacity: 30},
		{LowerLimit: 90, UpperLimit: 95, MaxCapacity: 0},
		{LowerLimit: 95, UpperLimit: 100, MaxCapacity: 30},
		{LowerLimit: 100, UpperLimit: 150, MaxCapacity: 20},
		{LowerLimit: 150, UpperLimit: 200, MaxCapacity: 20},
		{LowerLimit: 200, UpperLimit: 500, MaxCapacity: 10},
		{LowerLimit: 500, UpperLimit: 1000, MaxCapacity: 5},
		{LowerLimit: 1000, UpperLimit: 5000, MaxCapacity: 3},
	}
}
