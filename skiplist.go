package datastruct

import (
	"math/rand"
)

const (
	MAXLEVEL = 31
)

type node struct {
	value    interface{}
	score    int32
	backward *node
	forward  []*node
}

type SkipList struct {
	head  *node
	level int32
	count int32
}

func NewSkipList() *SkipList {
	return &SkipList{newSkipListNode(nil, 0, 32), 1, 0}
}

func newSkipListNode(value interface{}, score int32, level uint32) *node {
	return &node{value, score, nil, make([]*node, level)}
}

func randomLevel() int32 {

	level := 1

	for (rand.Int() & 0xFFFF) < (0.5 * 0xFFFF) {
		level += 1
	}

	if level > MAXLEVEL {
		return MAXLEVEL
	}

	return level
}

func (this *SkipList) Add(value interface{}, score int32) {

	var update [MAXLEVEL]*node
	var node *node

	node = this.head // start from head node

	for i := this.level - 1; i >= 0; i-- {
		for node.forward[i] != nil && node.forward[i].score < score {
			node = node.forward[i]
		}
		update[i] = node
	}

	level := randomLevel()

	if level > this.level {
		for i = this.level; i < level; i++ {
			update[i] = this.head
		}
		this.level = level
	}

	node = newSkipListNode(value, score, level)

	for i = 0; i < level; i++ {
		node.forward[i] = update[i].forward[i]
		update[i].forward[i] = node
	}

	if update[0] == this.head {
		node.backward = nil
	} else {
		node.backward = update[0]
	}

	if node.forward[0] != nil {
		node.forward[0].backward = node
	}

	this.count++
}
