package iterator

// Iterator interface contains the iterator functions.
type Iterator interface {
	HasNext() bool
	Next() any
}

// Iter keeps data and index.
type Iter struct {
	arr []any
	idx int
}

// New creates *Iter object by given elements in parameter. If nothing is passed
// it creates an empty slice.
func New(a ...any) *Iter {
	i := &Iter{}
	i.arr = append(i.arr, a...)
	return i
}

// HasNext returns whether iter has element.
func (i *Iter) HasNext() bool {
	return i.idx < len(i.arr)
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
