package gocache

import (
	"time"
)

var (
	DefaultCache = NewCache()
)

const (
	defaultCleanInterval = 1 * time.Second
)

type Item struct {
	Data     interface{}
	ExpireAt time.Time
}

func NewItem(data interface{}, dur time.Duration) *Item {
	return &Item{
		Data:     data,
		ExpireAt: time.Now().Add(dur),
	}
}

type Cache struct {
	tree   *Tree
	option *Option
}

func NewCache(opts ...Options) *Cache {
	c := &Cache{
		tree: &Tree{
			lock: &Lock{},
			root: NewNode(),
		},
		option: &Option{},
	}
	for _, op := range opts {
		op(c.option)
	}
	if c.option.cleanInterval == 0 {
		c.option.cleanInterval = defaultCleanInterval
	}
	return c
}

func (c *Cache) Get(key string) (value interface{}, exist bool) {
	v, exist := c.tree.Get(key)
	if !exist {
		return nil, false
	}
	item := v.(*Item)
	if item.ExpireAt.After(time.Now()) {
		return item.Data, true
	}
	return nil, false
}

func (c *Cache) Add(key string, value interface{}, dur time.Duration) {
	c.tree.Set(key, NewItem(value, dur))
}
