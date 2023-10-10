package list

import (
	"github.com/stretchr/testify/assert"
	"go-generic/errs"
	"testing"
)

func TestBuildDoubleLinkedList(t *testing.T) {
	testCases := []struct {
		name      string
		slice     []int
		wantSlice []int
	}{
		{
			name:      "[]",
			slice:     []int{},
			wantSlice: []int{},
		},
		{
			name:      "[1]",
			slice:     []int{1},
			wantSlice: []int{1},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			dll := BuildDoubleLinkedList[int](tc.slice)
			assert.Equal(t, tc.wantSlice, dll.ToSlice())
		})
	}
}

func TestDoubleLinkedList_Append(t *testing.T) {
	testCases := []struct {
		name      string
		dll       *DoubleLinkedList[int]
		index     int
		value     int
		wantVal   int
		wantErr   error
		wantSlice []int
	}{
		{
			name:      "append at index 0",
			dll:       BuildDoubleLinkedList[int]([]int{}),
			value:     5,
			wantErr:   nil,
			wantSlice: []int{5},
		},
		{
			name:      "append at index 1",
			dll:       BuildDoubleLinkedList[int]([]int{1}), //
			value:     5,
			wantErr:   nil,
			wantSlice: []int{1, 5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.dll.Append(tc.value)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantSlice, tc.dll.ToSlice())
		})
	}
}

func TestDoubleLinkedList_Get(t *testing.T) {
	testCases := []struct {
		name      string
		dll       *DoubleLinkedList[int]
		index     int
		wantErr   error
		wantVal   *Node[int]
		wantSlice []int
	}{
		{
			name:      "Get node at index -1",
			dll:       BuildDoubleLinkedList[int]([]int{1, 2}),
			index:     -1,
			wantErr:   errs.NewErrIndexOutOfRange(2, -1),
			wantVal:   nil,
			wantSlice: []int{1, 2},
		},
		{
			name:      "Get node at index 2",
			dll:       BuildDoubleLinkedList[int]([]int{1, 2}),
			index:     2,
			wantErr:   errs.NewErrIndexOutOfRange(2, 2),
			wantVal:   nil,
			wantSlice: []int{1, 2},
		},
		{
			name:      "Get node at index 1",
			dll:       BuildDoubleLinkedList[int]([]int{1, 2}),
			index:     1,
			wantErr:   nil,
			wantVal:   NewNode[int](2),
			wantSlice: []int{1, 2},
		},
		{
			name:      "Get node at index 0",
			dll:       BuildDoubleLinkedList[int]([]int{1, 2}),
			index:     0,
			wantErr:   nil,
			wantVal:   NewNode[int](1),
			wantSlice: []int{1, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			node, err := tc.dll.Get(tc.index)
			assert.Equal(t, tc.wantErr, err)
			if tc.wantVal != nil {
				assert.Equal(t, tc.wantVal.val, node.val)
			} else {
				assert.Equal(t, tc.wantVal, node)
			}

			assert.Equal(t, tc.wantSlice, tc.dll.ToSlice())
		})
	}
}

func TestDoubleLinkedList_IndexOf(t *testing.T) {
	testCases := []struct {
		name      string
		dll       *DoubleLinkedList[int]
		value     int
		wantVal   int
		wantSlice []int
	}{
		{
			name:      "indexof non-exist value 5",
			dll:       BuildDoubleLinkedList[int]([]int{1, 2}),
			value:     5,
			wantVal:   -1,
			wantSlice: []int{1, 2},
		},
		{
			name:      "indexof value 1",
			dll:       BuildDoubleLinkedList[int]([]int{1, 2}),
			value:     1,
			wantVal:   0,
			wantSlice: []int{1, 2},
		},
		{
			name:      "indexof value 2",
			dll:       BuildDoubleLinkedList[int]([]int{1, 2}),
			value:     2,
			wantVal:   1,
			wantSlice: []int{1, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			wantindex := tc.dll.IndexOf(tc.value)
			assert.Equal(t, tc.wantVal, wantindex)
			assert.Equal(t, tc.wantSlice, tc.dll.ToSlice())
		})
	}
}

func TestDoubleLinkedList_InsertAfter(t *testing.T) {
	testCases := []struct {
		name      string
		dll       *DoubleLinkedList[int]
		index     int
		value     int
		wantVal   int
		wantErr   error
		wantSlice []int
	}{
		{
			name:      "insert after index 0",
			dll:       BuildDoubleLinkedList[int]([]int{1, 2}),
			value:     5,
			wantErr:   nil,
			wantSlice: []int{1, 5, 2},
		},
		{
			name:      "insert after index 1",
			dll:       BuildDoubleLinkedList[int]([]int{1, 2}),
			value:     5,
			index:     1,
			wantErr:   nil,
			wantSlice: []int{1, 2, 5},
		},
		{
			name:      "insert after index 2",
			dll:       BuildDoubleLinkedList[int]([]int{1, 2}),
			value:     5,
			index:     2,
			wantErr:   errs.NewErrIndexOutOfRange(2, 2),
			wantSlice: []int{1, 2},
		},
		{
			name:      "insert after index -1",
			dll:       BuildDoubleLinkedList[int]([]int{1, 2}),
			value:     5,
			index:     -1,
			wantErr:   errs.NewErrIndexOutOfRange(2, -1),
			wantSlice: []int{1, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.dll.InsertAfter(tc.index, tc.value)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantSlice, tc.dll.ToSlice())
		})
	}
}

func TestDoubleLinkedList_InsertBefore(t *testing.T) {
	testCases := []struct {
		name      string
		dll       *DoubleLinkedList[int]
		index     int
		value     int
		wantVal   int
		wantErr   error
		wantSlice []int
	}{
		{
			name:      "insert before index 0",
			dll:       BuildDoubleLinkedList[int]([]int{1, 2}),
			value:     5,
			wantErr:   nil,
			wantSlice: []int{5, 1, 2},
		},
		{
			name:      "insert before index 1",
			dll:       BuildDoubleLinkedList[int]([]int{1, 2}),
			value:     5,
			index:     1,
			wantErr:   nil,
			wantSlice: []int{1, 5, 2},
		},
		{
			name:      "insert before index 2",
			dll:       BuildDoubleLinkedList[int]([]int{1, 2}),
			value:     5,
			index:     2,
			wantErr:   nil,
			wantSlice: []int{1, 2, 5},
		},
		{
			name:      "insert before index -1",
			dll:       BuildDoubleLinkedList[int]([]int{1, 2}),
			value:     5,
			index:     -1,
			wantErr:   errs.NewErrIndexOutOfRange(2, -1),
			wantSlice: []int{1, 2},
		},
		{
			name:      "insert before index 5",
			dll:       BuildDoubleLinkedList[int]([]int{1, 2}),
			value:     5,
			index:     5,
			wantErr:   errs.NewErrIndexOutOfRange(2, 5),
			wantSlice: []int{1, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.dll.InsertBefore(tc.index, tc.value)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantSlice, tc.dll.ToSlice())
		})
	}
}

func TestDoubleLinkedList_Remove(t *testing.T) {
	testCases := []struct {
		name      string
		dll       *DoubleLinkedList[int]
		value     int
		ok        bool
		wantVal   int
		wantSlice []int
	}{
		{
			name:      "remove non-exist value 5",
			dll:       BuildDoubleLinkedList[int]([]int{1, 2}),
			value:     5,
			ok:        false,
			wantVal:   -1,
			wantSlice: []int{1, 2},
		},
		{
			name:      "remove first element",
			dll:       BuildDoubleLinkedList[int]([]int{1, 2}),
			value:     1,
			ok:        true,
			wantVal:   0,
			wantSlice: []int{2},
		},
		{
			name:      "remove first element 1",
			dll:       BuildDoubleLinkedList[int]([]int{1, 2, 1}),
			value:     1,
			ok:        true,
			wantVal:   0,
			wantSlice: []int{2, 1},
		},
		{
			name:      "remove last element",
			dll:       BuildDoubleLinkedList[int]([]int{1, 2}),
			value:     2,
			ok:        true,
			wantVal:   1,
			wantSlice: []int{1},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val, ok := tc.dll.Remove(tc.value)
			assert.Equal(t, tc.ok, ok)
			assert.Equal(t, tc.wantVal, val)
			assert.Equal(t, tc.wantSlice, tc.dll.ToSlice())
		})
	}
}

func TestDoubleLinkedList_RemoveAt(t *testing.T) {
	testCases := []struct {
		name      string
		dll       *DoubleLinkedList[int]
		index     int
		wantErr   error
		wantVal   int
		wantSlice []int
	}{
		{
			name:      "remove value at index -1",
			dll:       BuildDoubleLinkedList[int]([]int{1, 2}),
			index:     -1,
			wantErr:   errs.NewErrIndexOutOfRange(2, -1),
			wantVal:   0,
			wantSlice: []int{1, 2},
		},
		{
			name:      "remove value at index 0",
			dll:       BuildDoubleLinkedList[int]([]int{1, 2}),
			index:     0,
			wantErr:   nil,
			wantVal:   1,
			wantSlice: []int{2},
		},
		{
			name:      "remove value at index 1",
			dll:       BuildDoubleLinkedList[int]([]int{1, 2}),
			index:     1,
			wantErr:   nil,
			wantVal:   2,
			wantSlice: []int{1},
		},
		{
			name:      "remove value at index 2",
			dll:       BuildDoubleLinkedList[int]([]int{1, 2}),
			index:     2,
			wantErr:   errs.NewErrIndexOutOfRange(2, 2),
			wantVal:   0,
			wantSlice: []int{1, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val, err := tc.dll.RemoveAt(tc.index)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantVal, val)
			assert.Equal(t, tc.wantSlice, tc.dll.ToSlice())
		})
	}
}
