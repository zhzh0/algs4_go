// multiset
// 单链表

package f

import "context"

type node struct {
	Item ItemType
	Next *node
}

type Bag struct {
	first *node
	n     int
}

func NewBag() *Bag {
	return &Bag{}
}

func (b *Bag) IsEmpty() bool {
	return b.first == nil
}

func (b *Bag) Size() int {
	return b.n
}

func (b *Bag) Add(item ItemType) {

	var old = b.first
	b.first = &node{item, old}
	b.n++
}

func (b *Bag) Iterator(cb IterableCallback) {
	var err error
	for i := b.first; i != nil; i = i.Next {
		if err = cb(i.Item); err != nil {
			break
		}
	}
}

// chan generator
func (b *Bag) Chan(ctx context.Context) <-chan ItemType {
	buf := make(chan ItemType)
	go func() {
		defer close(buf)
		for n := b.first; n != nil; n = n.Next {
			select {
			case <-ctx.Done():
				return // 通过ctx控制返回
			default:
				buf <- n.Item
			}
		}
	}()
	return buf
}
