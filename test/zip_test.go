package test

import (
	"reflect"
	"testing"

	"github.com/micotom/gfuncs"
)

func TestZip(t *testing.T) {

	tables := []struct {
		in  gfuncs.Tuple[[]float64, []int]
		exp []gfuncs.Tuple[float64, int]
	}{
		{
			in:  gfuncs.Tuple[[]float64, []int]{First: []float64{}, Second: []int{}},
			exp: []gfuncs.Tuple[float64, int]{},
		},
		{
			in:  gfuncs.Tuple[[]float64, []int]{First: []float64{1.0}, Second: []int{1}},
			exp: []gfuncs.Tuple[float64, int]{{First: 1.0, Second: 1}},
		},
		{
			in:  gfuncs.Tuple[[]float64, []int]{First: []float64{1.0, 2.0}, Second: []int{1}},
			exp: []gfuncs.Tuple[float64, int]{{First: 1.0, Second: 1}},
		},
		{
			in:  gfuncs.Tuple[[]float64, []int]{First: []float64{1.0}, Second: []int{1, 2}},
			exp: []gfuncs.Tuple[float64, int]{{First: 1.0, Second: 1}},
		},
	}

	for _, table := range tables {
		r := gfuncs.Zip(table.in.First, table.in.Second)
		if !reflect.DeepEqual(table.exp, r) {
			t.Errorf("Error on input %v: expected %v, got %v", table.in, table.exp, r)
		}
	}

}
