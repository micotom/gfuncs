package test

import (
	"reflect"
	"testing"

	"github.com/micotom/gfuncs"
)

func TestNone(t *testing.T) {
	tables := []struct {
		in  []int
		exp bool
	}{
		{
			in: []int{}, exp: true,
		},
		{
			in: []int{1}, exp: false,
		},
		{
			in: []int{2}, exp: true,
		},
		{
			in: []int{1, 2}, exp: false,
		},
	}

	fn := func(i int) bool {
		return i%2 != 0
	}
	for _, table := range tables {
		r := gfuncs.None(table.in, fn)
		if !reflect.DeepEqual(table.exp, r) {
			t.Errorf("Error on input %v: expected %v, got %v", table.in, table.exp, r)
		}
	}
}
