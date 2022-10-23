package test

import (
	"reflect"
	"testing"

	"github.com/micotom/gfuncs"
)

func TestFilter(t *testing.T) {
	tables := []struct {
		in  []int
		exp []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{}},
		{[]int{2}, []int{2}},
	}

	fn := func(i int) bool {
		return i%2 == 0
	}

	for _, table := range tables {
		r := gfuncs.Filter(table.in, fn)
		if !reflect.DeepEqual(table.exp, r) {
			t.Errorf("Error on input %v: expected %v, got %v", table.in, table.exp, r)
		}
	}
}

func TestFilterIndexed(t *testing.T) {
	tables := []struct {
		in  []int
		exp []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{}},
		{[]int{1, 4, 8}, []int{8}},
	}

	fn := func(i int, t int) bool {
		return i%2 == 0 && t%4 == 0
	}

	for _, table := range tables {
		r := gfuncs.FilterIndexed(table.in, fn)
		if !reflect.DeepEqual(table.exp, r) {
			t.Errorf("Error on input %v: expected %v, got %v", table.in, table.exp, r)
		}
	}
}
