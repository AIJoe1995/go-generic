package list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-generic/errs"
	"testing"
)

func TestArrayList_AddAt(t *testing.T) {
	testCases := []struct {
		name    string
		arr     *ArrayList[int]
		index   int
		value   int
		wantArr *ArrayList[int]
		wantErr error
	}{
		{
			name:    "add at index 0 to empty arraylist",
			arr:     NewArrayList[int](0),
			index:   0,
			value:   1,
			wantArr: NewArrayListFromSlice([]int{1}),
			wantErr: nil,
		},
		{
			name:    "add at index 0 to length2 arraylist",
			arr:     NewArrayListFromSlice([]int{1, 2}),
			index:   0,
			value:   0,
			wantArr: NewArrayListFromSlice([]int{0, 1, 2}),
			wantErr: nil,
		},
		{
			name:    "add at index 2 to length2 arraylist",
			arr:     NewArrayListFromSlice([]int{1, 2}),
			index:   2,
			value:   3,
			wantArr: NewArrayListFromSlice([]int{1, 2, 3}),
			wantErr: nil,
		},
		{
			name:    "add at index 1 to length2 arraylist",
			arr:     NewArrayListFromSlice([]int{1, 2}),
			index:   1,
			value:   5,
			wantArr: NewArrayListFromSlice([]int{1, 5, 2}),
			wantErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.arr.AddAt(tc.index, tc.value)
			assert.NoError(t, err)
			fmt.Println(tc.arr)
			assert.Equal(t, tc.wantArr, tc.arr)
		})
	}

}

func TestArrayList_Append(t *testing.T) {
	testCases := []struct {
		name    string
		arr     *ArrayList[int]
		value   int
		wantArr *ArrayList[int]
		wantErr error
	}{
		{
			name:    "append to empty arraylist",
			arr:     NewArrayList[int](0),
			value:   1,
			wantArr: NewArrayListFromSlice([]int{1}),
			wantErr: nil,
		},
		{
			name:    "append to length 2 arraylist",
			arr:     NewArrayListFromSlice([]int{1, 2}),
			value:   3,
			wantArr: NewArrayListFromSlice([]int{1, 2, 3}),
			wantErr: nil,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.arr.Append(tc.value)
			assert.NoError(t, err)
			assert.Equal(t, tc.wantArr, tc.arr)
		})
	}
}

func TestArrayList_DeleteAt(t *testing.T) {
	testCases := []struct {
		name    string
		arr     *ArrayList[int]
		index   int
		wantVal int
		wantArr *ArrayList[int]
		wantErr error
	}{
		{
			name:    "delete at index 0 to empty arraylist",
			arr:     NewArrayList[int](0),
			index:   0,
			wantVal: 0,
			wantArr: NewArrayListFromSlice([]int{}),
			wantErr: errs.NewErrIndexOutOfRange(0, 0),
		},
		{
			name:    "delete at index 3 to length 2 arraylist",
			arr:     NewArrayListFromSlice([]int{1, 2}),
			index:   3,
			wantVal: 0,
			wantArr: NewArrayListFromSlice([]int{1, 2}),
			wantErr: errs.NewErrIndexOutOfRange(2, 3),
		},
		{
			name:    "delete at index 0 to length 1 arraylist",
			arr:     NewArrayListFromSlice([]int{1}),
			index:   0,
			wantVal: 1,
			wantArr: NewArrayListFromSlice([]int{}),
			wantErr: nil,
		},
		{
			name:    "delete at index 1 to length 2 arraylist",
			arr:     NewArrayListFromSlice([]int{1, 2}),
			index:   1,
			wantVal: 2,
			wantArr: NewArrayListFromSlice([]int{1}),
			wantErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val, err := tc.arr.DeleteAt(tc.index)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantArr, tc.arr)
			assert.Equal(t, tc.wantVal, val)
		})
	}
}

func TestArrayList_Get(t *testing.T) {
	testCases := []struct {
		name    string
		arr     *ArrayList[int]
		index   int
		wantVal int
		wantArr *ArrayList[int]
		wantErr error
	}{
		{
			name:    "Get index 1 from length 1 array list",
			arr:     NewArrayListFromSlice[int]([]int{1}),
			index:   1,
			wantVal: 0,
			wantArr: NewArrayListFromSlice[int]([]int{1}),
			wantErr: errs.NewErrIndexOutOfRange(1, 1),
		},
		{
			name:    "Get index 1 from length 2 array list",
			arr:     NewArrayListFromSlice[int]([]int{1, 2}),
			index:   1,
			wantVal: 2,
			wantArr: NewArrayListFromSlice[int]([]int{1, 2}),
			wantErr: nil,
		},
		{
			name:    "Get index 3 from length 2 array list",
			arr:     NewArrayListFromSlice[int]([]int{1, 2}),
			index:   3,
			wantVal: 0,
			wantArr: NewArrayListFromSlice[int]([]int{1, 2}),
			wantErr: errs.NewErrIndexOutOfRange(2, 3),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			val, err := tc.arr.Get(tc.index)
			assert.Equal(t, tc.wantErr, err)
			assert.Equal(t, tc.wantVal, val)
			assert.Equal(t, tc.wantArr, tc.arr)
		})
	}

}

func TestArrayList_Set(t *testing.T) {
	testCases := []struct {
		name    string
		arr     *ArrayList[int]
		index   int
		value   int
		wantArr *ArrayList[int]
		wantErr error
	}{
		{
			name:    "set at index 0 to empty arraylist",
			arr:     NewArrayList[int](0),
			index:   0,
			value:   1,
			wantArr: NewArrayListFromSlice([]int{}),
			wantErr: errs.NewErrIndexOutOfRange(0, 0),
		},
		{
			name:    "set at index 1 to length2 arraylist",
			arr:     NewArrayListFromSlice([]int{1, 2}),
			index:   1,
			value:   5,
			wantArr: NewArrayListFromSlice([]int{1, 5}),
			wantErr: nil,
		},
		{
			name:    "add at index 2 to length2 arraylist",
			arr:     NewArrayListFromSlice([]int{1, 2}),
			index:   2,
			value:   3,
			wantArr: NewArrayListFromSlice([]int{1, 2}),
			wantErr: errs.NewErrIndexOutOfRange(2, 2),
		},
		{
			name:    "add at index -1 to length2 arraylist",
			arr:     NewArrayListFromSlice([]int{1, 2}),
			index:   -1,
			value:   5,
			wantArr: NewArrayListFromSlice([]int{1, 2}),
			wantErr: errs.NewErrIndexOutOfRange(2, -1),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.arr.Set(tc.index, tc.value)
			assert.Equal(t, tc.wantErr, err)
			//fmt.Println(tc.arr)
			assert.Equal(t, tc.wantArr, tc.arr)
		})
	}

}

//func TestNewArrayList(t *testing.T) {
//
//}
//
//func TestNewBArrayListFromSlice(t *testing.T) {
//
//}
