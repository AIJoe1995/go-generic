package slice

import (
	"github.com/stretchr/testify/assert"
	"go-generic/errs"
	"testing"
)

func TestAddAtIndex(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		element   int
		index     int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "add at index 0",
			slice:     []int{1, 2},
			element:   0,
			index:     0,
			wantSlice: []int{0, 1, 2},
			wantErr:   nil,
		},
		{
			name:      "add at index 1",
			slice:     []int{1, 3},
			element:   2,
			index:     1,
			wantSlice: []int{1, 2, 3},
			wantErr:   nil,
		},
		{
			name:      "index out of range, negative index",
			slice:     []int{1, 3},
			element:   4,
			index:     -1,
			wantSlice: nil,
			wantErr:   errs.NewErrIndexOutOfRange(2, -1),
		},
		{
			name:      "index out of range, greater than length",
			slice:     []int{1, 3},
			element:   4,
			index:     10,
			wantSlice: nil,
			wantErr:   errs.NewErrIndexOutOfRange(2, 10),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := AddAtIndex(tc.slice, tc.element, tc.index)
			assert.Equal(t, tc.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantSlice, res)
		})
	}
}
