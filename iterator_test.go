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
			name:     "Non-empty iterator, non-zero index",
			elements: []any{1, 2, 3, 4},
			idx:      2,
			hasNext:  true,
		},
		{
			name:     "Non-empty iterator, last index",
			elements: []any{1, 2, 3, 4},
			idx:      2, // TODO Check this part
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
			name:     "Non-empty iterator, non-zero index",
			elements: []any{1, 2, 3, 4},
			idx:      2,
			hasPrev:  true,
		},
		{
			name:     "Non-empty iterator, last index",
			elements: []any{1, 2, 3, 4},
			idx:      2, // TODO Check this part
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
			idx:      1, // TODO Check this part
			expVal:   2,
		},
		{
			name:     "Non-empty iterator, last index",
			elements: []any{1, 2, 3, 4, 5},
			idx:      3, // TODO Check this part
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
