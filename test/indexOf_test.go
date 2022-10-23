package test

import (
	"reflect"
	"testing"

	"github.com/micotom/gfuncs"
)

func TestIndexOf(t *testing.T) {
	tables := []struct {
		in  []int
		exp int
	}{
		{
			in: []int{}, exp: -1,
		},
		{
			in: []int{1}, exp: 0,
		},
		{
			in: []int{2, 1}, exp: 1,
		},
	}
	fn := func(i int) bool {
		return i == 1
	}
	for _, table := range tables {
		r := gfuncs.IndexOf(table.in, 1, fn)
		if !reflect.DeepEqual(table.exp, r) {
			t.Errorf("Error on input %v: expected %v, got %v", table.in, table.exp, r)
		}
	}
}
