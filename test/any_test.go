package test

import (
	"reflect"
	"testing"

	"github.com/micotom/gfuncs"
)

func TestAny(t *testing.T) {
	tables := []struct {
		in  []int
		exp bool
	}{
		{[]int{}, false},
		{[]int{1}, false},
		{[]int{2}, true},
		{[]int{1, 2}, true},
	}
	fn := func(i int) bool {
		return i%2 == 0
	}
	for _, table := range tables {
		r := gfuncs.Any(table.in, fn)
		if !reflect.DeepEqual(table.exp, r) {
			t.Errorf("Error on input %v: expected %v, got %v", table.in, table.exp, r)
		}
	}
}
