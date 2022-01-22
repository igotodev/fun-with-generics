package collection

import (
	"errors"
	"sort"
)

type Allowed interface {
	signedInt | unsignedInt | floats | string
}

type signedInt interface {
	int | int8 | int16 | int32 | int64
}

type unsignedInt interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type floats interface {
	float32 | float64
}

type List[T Allowed] struct {
	list []T
}

func NewList[T Allowed]() *List[T] {
	return new(List[T])
}

func (a *List[T]) Add(elem ...T) {
	a.list = append(a.list, elem...)
}

func (a *List[T]) Remove(index int) error {

	if len(a.list) <= index {
		return errors.New("error: index is larger than collection size")
	}
	a.list = append(a.list[:index], a.list[index+1:]...)
	return nil
}

func (a *List[T]) RemoveAll() {

	a.list = append(a.list[0:0])
}

func (a List[T]) Get(index int) (T, error) {
	var def T
	if len(a.list) <= index {
		return def, errors.New("error: invalid index")
	}
	return a.list[index], nil
}

func (a List[T]) GetAll() []T {

	return a.list
}
func (a *List[T]) Sort() {
	sort.Sort(a)
}

func (a *List[T]) Len() int {
	return len(a.list)
}

func (a *List[T]) IsEmpty() bool {
	if len(a.list) == 0 {
		return true
	}
	return false
}

// Swap swaps the elements with indexes i and j.
func (a *List[T]) Swap(i, j int) {
	mid := a.list[j]
	a.list[j] = a.list[i]
	a.list[i] = mid
}

// Less return true if element with index i less than element with index j.
func (a *List[T]) Less(i, j int) bool {
	if a.list[i] < a.list[j] {
		return true
	}
	return false
}
