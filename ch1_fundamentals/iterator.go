package f

type ItemType int // 泛型的忧伤

type IterableCallback func(ItemType) error

type Iterable interface {
	Iterator(IterableCallback)
}
