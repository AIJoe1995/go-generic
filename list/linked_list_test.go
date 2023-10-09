package list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceBuild(t *testing.T) {
	slice := []int{1, 2}
	fmt.Printf("%T", slice)
}

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

}

func TestDoubleLinkedList_IndexOf(t *testing.T) {

}

func TestDoubleLinkedList_InsertAfter(t *testing.T) {

}

func TestDoubleLinkedList_InsertBefore(t *testing.T) {

}

func TestDoubleLinkedList_Remove(t *testing.T) {

}

func TestDoubleLinkedList_RemoveAt(t *testing.T) {

}
