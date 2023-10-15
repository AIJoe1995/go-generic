package queue

import (
	"github.com/stretchr/testify/assert"
	"go-generic/errs"
	"testing"
)

func TestSimpleQueue_SliceToSimpleQueue(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		wantSlice []int
		wantErr   error
	}{
		{
			name:      "[1,2,3] slice to queue back to slice",
			slice:     []int{1, 2, 3},
			wantSlice: QueueToSlice[int](SliceToSimpleQueue[int]([]int{1, 2, 3})),
			wantErr:   nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.slice, tc.wantSlice)
		})
	}
}

func TestSimpleQueue_Dequeue(t *testing.T) {
	testCases := []struct {
		name    string
		queue   *SimpleQueue[int]
		wantErr error
		wantVal int
		wantLen int
	}{
		{
			name:    "[1,2,3] slice to queue ",
			queue:   SliceToSimpleQueue[int]([]int{1, 2, 3}),
			wantErr: nil,
			wantVal: 1,
			wantLen: 2,
		},
		{
			name:    "[] slice to queue ",
			queue:   NewSimpleQueue[int](),
			wantErr: errs.NewErrEmptyQueue(),
			wantVal: 0,
			wantLen: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val, err := tc.queue.Dequeue()
			assert.Equal(t, tc.wantLen, tc.queue.Len())
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantVal, val)

		})
	}
}

//func TestSimpleQueue_Enqueue(t *testing.T) {
//
//}

func TestSimpleQueue_Peek(t *testing.T) {

}
