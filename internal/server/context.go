package server

import (
	"encoding/json"
	"game_server_slots_fortune_snake/internal/model"
	"sync"
)

type Context struct {
	engine   *Engine
	handlers []HandlerFunc
	mu       sync.RWMutex
	keys     map[string]interface{}
	index    int
	data     []byte
	result   []byte
}

func (c *Context) SendError(code int32, msg string) {
	resp := &model.MessageResponse{
		Code:    code,
		Message: msg,
	}
	c.result, _ = json.Marshal(resp)
}

func (c *Context) Send(code int32, msg string, data interface{}) {
	jsonData, _ := json.Marshal(data)
	resp := &model.MessageResponse{
		Code:    code,
		Message: msg,
		Data:    string(jsonData),
	}
	c.result, _ = json.Marshal(resp)
}

func (c *Context) Data() []byte {
	return c.data
}

func (c *Context) Set(key string, value interface{}) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	c.keys[key] = value
}

func (c *Context) Get(key string) (value interface{}, exists bool) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	value, exists = c.keys[key]
	return
}

func (c *Context) MustGet(key string) interface{} {
	if value, exists := c.Get(key); exists {
		return value
	}
	panic("Key \"" + key + "\" does not exist")
}

func (c *Context) Next() {
	c.index++
	for c.index < len(c.handlers) {
		c.handlers[c.index](c)
		c.index++
	}
}

func (c *Context) Abort() {
	c.index = len(c.handlers)
}
