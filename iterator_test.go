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
			idx:      2,
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
