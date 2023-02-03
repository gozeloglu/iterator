package iterator

// Iterator interface contains the iterator functions.
type Iterator interface {
	HasNext() bool
	HasPrev() bool
	Next() any
	Prev() any
	ToSlice() []any
}

// Iter keeps data and index.
type Iter struct {
	arr []any
	idx int
}

// New creates *Iter object by given elements in parameter. If nothing is passed
// it creates an empty slice.
func New(a ...any) *Iter {
	var arr []any
	i := &Iter{
		arr: arr,
	}
	i.arr = append(i.arr, a...)
	return i
}

// HasNext returns whether iter has element in next.
func (i *Iter) HasNext() bool {
	return i.idx < len(i.arr)
}

// HasPrev returns whether iter has element in previous.
func (i *Iter) HasPrev() bool {
	return i.idx >= 0 && i.idx <= len(i.arr) && len(i.arr) != 0
}

// Next returns the next element. If there is no next element, it returns nil.
func (i *Iter) Next() any {
	if i.HasNext() {
		if i.idx < 0 {
			i.idx = 0
		}
		v := i.arr[i.idx]
		i.idx++
		return v
	}
	return nil
}

// Prev returns previous element. If there is no previous element, it returns nil.
func (i *Iter) Prev() any {
	if i.HasPrev() {
		if i.idx == len(i.arr) {
			i.idx--
		}
		v := i.arr[i.idx]
		i.idx--
		return v
	}
	return nil
}

// ToSlice returns data.
func (i *Iter) ToSlice() []any {
	return i.arr
}
