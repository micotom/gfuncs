package test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/micotom/gfuncs"
)

func TestMap(t *testing.T) {

	tables := []struct {
		in  []int
		exp []string
	}{
		{[]int{}, []string{}},
		{[]int{1}, []string{"one"}},
		{[]int{1, 2}, []string{"one", "two"}},
	}

	fn := func(i int) string {
		switch i {
		case 1:
			return "one"
		case 2:
			return "two"
		}
		panic("oops!")
	}

	for _, table := range tables {
		r := gfuncs.Map(table.in, fn)
		if !reflect.DeepEqual(table.exp, r) {
			t.Errorf("Error on input %v: expected %v, got %v", table.in, table.exp, r)
		}
	}

}

func TestMapIndexed(t *testing.T) {

	tables := []struct {
		in  []int
		exp []string
	}{
		{[]int{}, []string{}},
		{[]int{1}, []string{"one: 0"}},
		{[]int{1, 2}, []string{"one: 0", "two: 1"}},
	}

	fn := func(i int, t int) string {
		switch t {
		case 1:
			return fmt.Sprintf("one: %d", i)
		case 2:
			return fmt.Sprintf("two: %d", i)
		}
		panic("oops!")
	}

	for _, table := range tables {
		r := gfuncs.MapIndexed(table.in, fn)
		if !reflect.DeepEqual(table.exp, r) {
			t.Errorf("Error on input %v: expected %v, got %v", table.in, table.exp, r)
		}
	}

}
