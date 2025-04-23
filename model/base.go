package model

import "context"

type BaseInput struct {
	Ctx context.Context
}

type BaseOutput struct {
	Code    int32
	Message string
}
