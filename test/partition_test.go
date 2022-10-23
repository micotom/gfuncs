package test

import (
	"reflect"
	"testing"

	"github.com/micotom/gfuncs"
)

func TestPartition(t *testing.T) {
	tables := []struct {
		in  []int
		exp [][]int
	}{
		{
			in: []int{}, exp: [][]int{{}, {}},
		},
		{
			in: []int{1}, exp: [][]int{{}, {1}},
		},
		{
			in: []int{1, 2}, exp: [][]int{{2}, {1}},
		},
	}
	fn := func(i int) bool {
		return i%2 == 0
	}
	for _, table := range tables {
		r := gfuncs.Partition(table.in, fn)
		if !reflect.DeepEqual(table.exp, r) {
			t.Errorf("Error on input %v: expected %v, got %v", table.in, table.exp, r)
		}
	}
}
