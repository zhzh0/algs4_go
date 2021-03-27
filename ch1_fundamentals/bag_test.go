package f

import (
	"context"
	"testing"
)

func TestBag(t *testing.T) {
	b := NewBag()
	b.Add(ItemType(1))
	b.Add(ItemType(2))
	b.Add(ItemType(1))
	b.Add(ItemType(3))
	b.Add(ItemType(4))
	if b.Size() != 5 {
		t.Error("size not match", b.Size())
	}

	// iterator
	//print
	b.Iterator(func(item ItemType) error {
		t.Log(item)
		return nil
	})
	// sum
	total := 0
	b.Iterator(func(item ItemType) error {
		total += int(item)
		return nil
	})
	t.Logf("total is %d", total)

	// generator
	for item := range b.Chan(context.Background()) {
		t.Log(item)
	}
}

func BenchmarkBag(b *testing.B) {
	bag := NewBag()
	for i := 0; i < 10000000; i++ {
		bag.Add(ItemType(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _ = range bag.Chan(context.TODO()) {

		}
	}
}

func BenchmarkBag1(b *testing.B) {
	bag := NewBag()
	for i := 0; i < 10000000; i++ {
		bag.Add(ItemType(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bag.Iterator(func(item ItemType) error {
			return nil
		})
	}
}
