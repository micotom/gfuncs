package test

import (
	"reflect"
	"testing"

	"github.com/micotom/gfuncs"
)

func TestFind(t *testing.T) {
	var target *int
	target = new(int)
	*target = 42
	tables := []struct {
		in  []int
		exp *int
	}{
		{[]int{}, nil},
		{[]int{*target}, target},
		{[]int{1, 2, 3}, nil},
	}
	fn := func(i int) bool {
		return i == *target
	}
	for _, table := range tables {
		r := gfuncs.Find(table.in, fn)
		if !reflect.DeepEqual(table.exp, r) {
			t.Errorf("Error on input %v: expected %v, got %v", table.in, table.exp, r)
		}
	}
}
