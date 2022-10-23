package test

import (
	"reflect"
	"testing"

	"github.com/micotom/gfuncs"
)

func TestFold(t *testing.T) {
	tables := []struct {
		in  []string
		exp string
	}{
		{
			[]string{}, "",
		},
		{
			[]string{"one"}, "one",
		},
		{
			[]string{"one", "two"}, "onetwo",
		},
	}

	fn := func(acc string, s string) string {
		return acc + s
	}

	initial := new(string)
	for _, table := range tables {
		r := gfuncs.Fold(table.in, *initial, fn)
		if !reflect.DeepEqual(table.exp, r) {
			t.Errorf("Error on input %v: expected %v, got %v", table.in, table.exp, r)
		}
	}
}
