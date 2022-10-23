package test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/micotom/gfuncs"
)

type sumBymockStruct struct {
	Id    string
	Value int
}

type ErrorOrResult[T gfuncs.Number] struct {
	result *T
	err    error
}

func TestSumBy(t *testing.T) {
	tables := []struct {
		in  []sumBymockStruct
		exp ErrorOrResult[int]
	}{
		{[]sumBymockStruct{}, ErrorOrResult[int]{nil, errors.New("Cannot sum slice of length 0")}},
		{[]sumBymockStruct{{"1", 1}}, ErrorOrResult[int]{createIntPointer(1), nil}},
		{[]sumBymockStruct{{"1", 1}, {"2", 1}}, ErrorOrResult[int]{createIntPointer(2), nil}},
	}

	fn := func(m sumBymockStruct) int {
		return m.Value
	}

	for _, table := range tables {
		rTemp, err := gfuncs.SumBy(table.in, fn)
		r := ErrorOrResult[int]{rTemp, err}

		if !reflect.DeepEqual(table.exp, r) {
			t.Errorf("Error on input %v: expected %v, got %v", table.in, table.exp, r)
		}
	}

}

func createIntPointer(v int) *int {
	i := new(int)
	*i = v
	return i
}
