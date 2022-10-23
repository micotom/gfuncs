package test

import (
	"reflect"
	"testing"

	"github.com/micotom/gfuncs"
)

func TestSortyBy(t *testing.T) {
	tables := []struct {
		in  []int
		exp []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{1, 2, 1}, []int{1, 1, 2}},
	}

	fn := func(i int) int {
		return i
	}

	for _, table := range tables {
		r := gfuncs.SortBy(table.in, fn)
		if !reflect.DeepEqual(table.exp, r) {
			t.Errorf("Error on input %v: expected %v, got %v", table.in, table.exp, r)
		}
	}
}
