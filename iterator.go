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
	i := &Iter{
		arr: make([]any, len(a)),
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

// Next returns next element.
func (i *Iter) Next() any {
	if i.HasNext() {
		v := i.arr[i.idx]
		i.idx++
		return v
	}
	return nil
}

// Prev returns previous element.
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
