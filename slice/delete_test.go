package slice

import (
	"github.com/stretchr/testify/assert"
	"go-generic/errs"
	"testing"
)

func TestDeleteAtIndex(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		index     int
		wantSlice []int
		wantElem  int
		wantErr   error
	}{

		{
			name:      "delete at index 0",
			slice:     []int{1, 2},
			index:     0,
			wantSlice: []int{2},
			wantElem:  1,
			wantErr:   nil,
		},
		{
			name:      "delete at index 0, length 1",
			slice:     []int{1},
			index:     0,
			wantSlice: []int{},
			wantElem:  1,
			wantErr:   nil,
		},
		{
			name:      "index out of range, negative index",
			slice:     []int{1, 3},
			index:     -1,
			wantSlice: nil,
			wantElem:  0,
			wantErr:   errs.NewErrIndexOutOfRange(2, -1),
		},
		{
			name:      "index out of range, greater than length",
			slice:     []int{1, 3},
			index:     10,
			wantSlice: nil,
			wantElem:  0,
			wantErr:   errs.NewErrIndexOutOfRange(2, 10),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, elem, err := DeleteAtIndex(tc.slice, tc.index)
			assert.Equal(t, tc.wantSlice, res)
			if err != nil {
				return
			}
			assert.Equal(t, tc.wantElem, elem)
			assert.Equal(t, tc.wantSlice, res)
		})
	}

}
