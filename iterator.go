package iterator

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
		idx: -1,
	}
	i.arr = append(i.arr, a...)
	return i
}

// Next returns whether iter has element in next. If there is no next element,
// it returns false.
func (i *Iter) Next() bool {
	if i.idx < -1 {
		i.idx = -1
	}
	i.idx++
	return i.hasNext()
}

// Prev returns whether iter has element in prev. If there is no previous
// element, it returns false.
func (i *Iter) Prev() bool {
	if i.idx > len(i.arr) {
		i.idx = len(i.arr)
	}
	i.idx--
	return i.hasPrev()
}

// Value returns current value. If there is no value, it returns nil.
func (i *Iter) Value() any {
	if i.idx < 0 || i.idx >= len(i.arr) {
		return nil
	}
	return i.arr[i.idx]
}

// First updates the cursor by moving to first index.
func (i *Iter) First() {
	i.idx = -1
}

// Last updates the cursor by moving to last index. If array is empty, the cursor
// will be zero.
func (i *Iter) Last() {
	if len(i.arr) == 0 {
		i.idx = 0
		return
	}
	i.idx = len(i.arr)
}

// ToSlice returns data.
func (i *Iter) ToSlice() []any {
	return i.arr
}

// hasNext returns whether iter has element in next.
func (i *Iter) hasNext() bool {
	return i.idx < len(i.arr)
}

// hasPrev returns whether iter has element in previous.
func (i *Iter) hasPrev() bool {
	return (i.idx+1) > 0 && (i.idx+1) <= len(i.arr) && len(i.arr) != 0
}
