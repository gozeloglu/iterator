package iterator

import "testing"

func TestIter_HasNext(t *testing.T) {
	testCases := []struct {
		name     string
		elements []any
		idx      int
		hasNext  bool
	}{
		{
			name:     "Empty iterator",
			elements: []any{},
		},
		{
			name:     "Non-empty iterator, -1 index",
			elements: []any{1, 2, 3},
			idx:      -1,
			hasNext:  true,
		},
		{
			name:     "Non-empty iterator, fresh index",
			elements: []any{1, 2, 3},
			hasNext:  true,
		},
		{
			name:     "Non-empty iterator, 1-length",
			elements: []any{1},
			hasNext:  true,
		},
		{
			name:     "Non-empty iterator, non-zero index",
			elements: []any{1, 2, 3, 4},
			idx:      2,
			hasNext:  true,
		},
		{
			name:     "Non-empty iterator, last index",
			elements: []any{1, 2, 3, 4},
			idx:      3,
			hasNext:  true,
		},
		{
			name:     "Non-empty iterator, out of index",
			elements: []any{1, 2, 3, 4},
			idx:      4,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			it := New(tc.elements...)
			it.idx = tc.idx
			ok := it.HasNext()
			if tc.hasNext != ok {
				t.Fatalf("expected: %v\nactual: %v\n", tc.hasNext, ok)
			}
		})
	}
}

func TestIter_Next(t *testing.T) {
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
			expVal:   1,
		},
		{
			name:     "Non-empty iterator, fresh index",
			elements: []any{1, 2, 3, 4, 5},
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
			idx:      5,
			expVal:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			it := New(tc.elements...)
			it.idx = tc.idx
			v := it.Next()
			if tc.expVal != v {
				t.Fatalf("expected: %v\nactual: %v\n", tc.expVal, v)
			}
		})
	}
}

func TestIter_HasPrev(t *testing.T) {
	testCases := []struct {
		name     string
		elements []any
		idx      int
		hasPrev  bool
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
		},
		{
			name:     "Non-empty iterator, non-zero index",
			elements: []any{1, 2, 3, 4},
			idx:      2,
			hasPrev:  true,
		},
		{
			name:     "Non-empty iterator, last index",
			elements: []any{1, 2, 3, 4},
			idx:      3,
			hasPrev:  true,
		},
		{
			name:     "Non-empty iterator, out of index",
			elements: []any{1, 2, 3, 4},
			idx:      4,
			hasPrev:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			it := New(tc.elements...)
			it.idx = tc.idx
			ok := it.HasPrev()
			if tc.hasPrev != ok {
				t.Fatalf("expected: %v\nactual: %v\n", tc.hasPrev, ok)
			}
		})
	}
}

func TestIter_Prev(t *testing.T) {
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
		},
		{
			name:     "Non-empty iterator, fresh index",
			elements: []any{1, 2, 3, 4, 5},
		},
		{
			name:     "Non-empty iterator, non-zero index",
			elements: []any{1, 2, 3, 4, 5},
			idx:      1,
			expVal:   2,
		},
		{
			name:     "Non-empty iterator, last index",
			elements: []any{1, 2, 3, 4, 5},
			idx:      3,
			expVal:   4,
		},
		{
			name:     "Non-empty iterator, out of index",
			elements: []any{1, 2, 3, 4, 5},
			idx:      5,
			expVal:   5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			it := New(tc.elements...)
			it.idx = tc.idx
			v := it.Prev()
			if tc.expVal != v {
				t.Fatalf("expected: %v\nactual: %v\n", tc.expVal, v)
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
			expIdx := 0
			it := New(tc.elements...)
			it.idx = tc.idx
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
			expIdx:   2,
		},
		{
			name:     "Non-empty iterator, fresh index",
			elements: []any{1, 2, 3},
			idx:      0,
			expIdx:   2,
		},
		{
			name:     "Non-empty iterator, 1-length",
			elements: []any{1},
			idx:      0,
			expIdx:   0,
		},
		{
			name:     "Non-empty iterator, non-zero index",
			elements: []any{1, 2, 3, 4},
			idx:      2,
			expIdx:   3,
		},
		{
			name:     "Non-empty iterator, last index",
			elements: []any{1, 2, 3, 4},
			idx:      3,
			expIdx:   3,
		},
		{
			name:     "Non-empty iterator, out of index",
			elements: []any{1, 2, 3, 4},
			idx:      4,
			expIdx:   3,
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
