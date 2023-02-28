package iterator

import "testing"

func TestIter_Next(t *testing.T) {
	testCases := []struct {
		name     string
		elements []any
		idx      int
		expBool  bool
		expIdx   int
	}{
		{
			name:     "Empty iterator",
			elements: []any{},
			expBool:  false,
			expIdx:   0,
		},
		{
			name:     "Non-empty iterator, -2 index",
			elements: []any{1, 2, 3, 4, 5},
			idx:      -2,
			expBool:  true,
			expIdx:   0,
		},
		{
			name:     "Non-empty iterator, fresh index",
			elements: []any{1, 2, 3, 4, 5},
			expBool:  true,
			expIdx:   0,
		},
		{
			name:     "Non-empty iterator, non-zero index",
			elements: []any{1, 2, 3, 4, 5},
			idx:      2,
			expBool:  true,
			expIdx:   3,
		},
		{
			name:     "Non-empty iterator, last index",
			elements: []any{1, 2, 3, 4, 5},
			idx:      4,
			expBool:  false,
			expIdx:   5,
		},
		{
			name:     "Non-empty iterator, out of index",
			elements: []any{1, 2, 3, 4, 5},
			idx:      5,
			expBool:  false,
			expIdx:   6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			it := New(tc.elements...)
			if tc.idx != 0 {
				it.idx = tc.idx
			}
			ok := it.Next()
			if tc.expBool != ok {
				t.Fatalf("expected: %v\nactual: %v\n", tc.expBool, ok)
			}
			if tc.expIdx != it.idx {
				t.Fatalf("expected: %d\nactual: %d\n", tc.expIdx, it.idx)
			}
		})
	}
}

func TestIter_Prev(t *testing.T) {
	testCases := []struct {
		name     string
		elements []any
		idx      int
		expBool  bool
		expIdx   int
	}{
		{
			name:     "Empty iterator",
			elements: []any{},
			expBool:  false,
			expIdx:   -2,
		},
		{
			name:     "Non-empty iterator,  index",
			elements: []any{1, 2, 3, 4, 5},
			idx:      -2,
			expBool:  false,
			expIdx:   -3,
		},
		{
			name:     "Non-empty iterator, fresh index",
			elements: []any{1, 2, 3, 4, 5},
			expBool:  false,
			expIdx:   -2,
		},
		{
			name:     "Non-empty iterator, non-zero index",
			elements: []any{1, 2, 3, 4, 5},
			idx:      1,
			expBool:  true,
			expIdx:   0,
		},
		{
			name:     "Non-empty iterator, last index",
			elements: []any{1, 2, 3, 4, 5},
			idx:      4,
			expBool:  true,
			expIdx:   3,
		},
		{
			name:     "Non-empty iterator, out of index",
			elements: []any{1, 2, 3, 4, 5},
			idx:      6,
			expBool:  true,
			expIdx:   4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			it := New(tc.elements...)
			if tc.idx != 0 {
				it.idx = tc.idx
			}
			ok := it.Prev()
			if tc.expBool != ok {
				t.Fatalf("expected: %v\nactual: %v\n", tc.expBool, ok)
			}
			if tc.expIdx != it.idx {
				t.Fatalf("expected: %d\nactual: %d\n", tc.expIdx, it.idx)
			}
		})
	}
}

func TestIter_First(t *testing.T) {
	testCases := []struct {
		name     string
		elements []any
		idx      int
	}{
		{
			name:     "Empty iterator",
			elements: []any{},
		},
		{
			name:     "Non-empty iterator, -1 index",
			elements: []any{1, 2, 3},
			idx:      -1,
		},
		{
			name:     "Non-empty iterator, fresh index",
			elements: []any{1, 2, 3},
			idx:      0,
		},
		{
			name:     "Non-empty iterator, 1-length",
			elements: []any{1},
			idx:      0,
		},
		{
			name:     "Non-empty iterator, non-zero index",
			elements: []any{1, 2, 3, 4},
			idx:      2,
		},
		{
			name:     "Non-empty iterator, last index",
			elements: []any{1, 2, 3, 4},
			idx:      3,
		},
		{
			name:     "Non-empty iterator, out of index",
			elements: []any{1, 2, 3, 4},
			idx:      4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expIdx := -1
			it := New(tc.elements...)
			if tc.idx != 0 {
				it.idx = tc.idx
			}
			it.First()
			if it.idx != expIdx {
				t.Errorf("expected idx: %d\nactual idx: %d\n", expIdx, it.idx)
			}
		})
	}
}

func TestIter_Last(t *testing.T) {
	testCases := []struct {
		name     string
		elements []any
		idx      int
		expIdx   int
	}{
		{
			name:     "Empty iterator",
			elements: []any{},
		},
		{
			name:     "Non-empty iterator, -1 index",
			elements: []any{1, 2, 3},
			idx:      -1,
			expIdx:   3,
		},
		{
			name:     "Non-empty iterator, fresh index",
			elements: []any{1, 2, 3},
			idx:      0,
			expIdx:   3,
		},
		{
			name:     "Non-empty iterator, 1-length",
			elements: []any{1},
			idx:      0,
			expIdx:   1,
		},
		{
			name:     "Non-empty iterator, non-zero index",
			elements: []any{1, 2, 3, 4},
			idx:      2,
			expIdx:   4,
		},
		{
			name:     "Non-empty iterator, last index",
			elements: []any{1, 2, 3, 4},
			idx:      3,
			expIdx:   4,
		},
		{
			name:     "Non-empty iterator, out of index",
			elements: []any{1, 2, 3, 4},
			idx:      5,
			expIdx:   4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			it := New(tc.elements...)
			it.idx = tc.idx
			it.Last()
			if it.idx != tc.expIdx {
				t.Errorf("expected idx: %d\nactual idx: %d\n", tc.expIdx, it.idx)
			}
		})
	}
}

func TestIter_ToSlice(t *testing.T) {
	testCases := []struct {
		name     string
		elements []any
	}{
		{
			name:     "Empty iterator",
			elements: []any{},
		},
		{
			name:     "Non-empty iterator",
			elements: []any{1, 2, 3, 4},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			it := New(tc.elements...)
			got := it.ToSlice()
			if len(tc.elements) != len(got) {
				t.Fatalf("expected length: %d\nactual lenght: %d\n", len(tc.elements), len(got))
			}
			for i, e := range got {
				if e != tc.elements[i] {
					t.Fatalf("expected element: %v\nactual element: %v\n", tc.elements[i], e)
				}
			}
		})
	}
}

func TestIter_Value(t *testing.T) {
	testCases := []struct {
		name     string
		elements []any
		idx      int
		expVal   any
	}{
		{
			name:     "Empty iterator",
			elements: []any{},
			expVal:   nil,
		},
		{
			name:     "Non-empty iterator, -1 index",
			elements: []any{1, 2, 3, 4, 5},
			idx:      -1,
			expVal:   nil,
		},
		{
			name:     "Non-empty iterator, 0 index",
			elements: []any{1, 2, 3, 4, 5},
			idx:      0,
			expVal:   1,
		},
		{
			name:     "Non-empty iterator, non-zero index",
			elements: []any{1, 2, 3, 4, 5},
			idx:      2,
			expVal:   3,
		},
		{
			name:     "Non-empty iterator, last index",
			elements: []any{1, 2, 3, 4, 5},
			idx:      4,
			expVal:   5,
		},
		{
			name:     "Non-empty iterator, out of index",
			elements: []any{1, 2, 3, 4, 5},
			idx:      6,
			expVal:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			it := New(tc.elements...)
			if tc.idx != -1 {
				it.idx = tc.idx
			}
			v := it.Value()
			if tc.expVal != v {
				t.Fatalf("expected value: %v\nactual value: %v\n", tc.expVal, v)
			}
		})
	}
}
