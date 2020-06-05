package gocache

import (
	"sync/atomic"
)

type Node struct {
	value atomic.Value
	child map[int32]*Node
	lock  *Lock
}

func NewNode() *Node {
	return &Node{
		value: atomic.Value{},
		child: make(map[int32]*Node),
		lock:  &Lock{},
	}
}

type Tree struct {
	count int32
	root  *Node
	lock  *Lock
}

func (tree *Tree) Get(key string) (interface{}, bool) {
	cur := tree.root
	for _, ch := range key {
		cur.lock.BeginRead()
		node, exist := cur.child[ch]
		cur.lock.EndRead()
		if !exist {
			return nil, false
		}
		cur = node
	}
	v := cur.value.Load()
	if v == nil {
		return nil, false
	} else {
		return v, true
	}
}

func (tree *Tree) Set(key string, value interface{}) {
	if key == "" || value == nil {
		return
	}
	cur := tree.root
	for _, ch := range key {
		cur.lock.BeginRead()
		node, exist := cur.child[ch]
		cur.lock.EndRead()
		if exist {
			cur = node
		} else {
			cur.lock.BeginWrite()
			temp := NewNode()
			cur.child[ch] = temp
			cur.lock.EndWrite()
			cur = temp
		}
	}
	cur.value.Store(value)
	return
}
