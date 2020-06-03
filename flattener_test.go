package flatten

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	flattenTests := []struct {
		twoD [][]int
		oneD []int
	}{
		{
			twoD: [][]int{{1}},
			oneD: []int{1},
		},
		{
			twoD: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			oneD: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			twoD: [][]int{
				{1, 2, 3, 1},
				{4, 5, 6, 4},
				{7, 8, 9, 7},
				{7, 8, 9, 7},
			},
			oneD: []int{1, 2, 3, 1, 4, 7, 7, 9, 8, 7, 7, 4, 5, 6, 9, 8},
		},
	}

	for _, tc := range flattenTests {
		init := tc.twoD
		want := tc.oneD
		got := Flatten(tc.twoD)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v to be converted to %v, instead got %v", init, want, got)
		}
	}
}

func ExampleFlatten() {
	flattened := Flatten([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	fmt.Println(flattened)
	// Output: [1 2 3 6 9 8 7 4 5]
}
