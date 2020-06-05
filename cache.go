package gocache

import (
	"sync"
	"time"
)

type Item struct {
	Data interface{}

	CreateTime time.Time

	Duration time.Duration
}

func (i *Item) GetData() interface{} {
	return i.Data
}

func NewItem(data interface{}, dur time.Duration) *Item {
	return &Item{
		Data:       data,
		CreateTime: time.Now(),
		Duration:   dur,
	}
}

type Cache struct {
	cache         map[interface{}]*Item
	addedItem     []func(item *Item)
	itemPool      sync.Pool
	cleanInterval time.Duration
}
