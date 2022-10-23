package test

import (
	"reflect"
	"testing"

	"github.com/micotom/gfuncs"
)

func TestFlatten(t *testing.T) {
	tables := []struct {
		in  [][]int
		exp []int
	}{
		{
			[][]int{}, []int{},
		},
		{
			[][]int{{1}}, []int{1},
		},
		{
			[][]int{{1}, {1}}, []int{1, 1},
		},
	}

	for _, table := range tables {
		r := gfuncs.Flatten(table.in)
		if !reflect.DeepEqual(table.exp, r) {
			t.Errorf("Error on input %v: expected %v, got %v", table.in, table.exp, r)
		}
	}
}
