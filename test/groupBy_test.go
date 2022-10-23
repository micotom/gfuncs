package test

import (
	"reflect"
	"testing"

	"github.com/micotom/gfuncs"
)

type groupByMockStruct struct {
	Id   int
	Name string
}

func TestGroupBy(t *testing.T) {

	tables := []struct {
		in  []groupByMockStruct
		exp map[string][]groupByMockStruct
	}{
		{[]groupByMockStruct{}, map[string][]groupByMockStruct{}},
		{
			[]groupByMockStruct{
				{Id: 1, Name: "1"},
			},
			map[string][]groupByMockStruct{
				"1": {groupByMockStruct{Id: 1, Name: "1"}},
			},
		},
		{
			[]groupByMockStruct{
				{Id: 1, Name: "1"},
				{Id: 2, Name: "1"},
			},
			map[string][]groupByMockStruct{
				"1": {
					groupByMockStruct{Id: 1, Name: "1"},
					groupByMockStruct{Id: 2, Name: "1"},
				},
			},
		},
		{
			[]groupByMockStruct{
				{Id: 1, Name: "1"},
				{Id: 2, Name: "2"},
			},
			map[string][]groupByMockStruct{
				"1": {
					groupByMockStruct{Id: 1, Name: "1"},
				},
				"2": {
					groupByMockStruct{Id: 2, Name: "2"},
				},
			},
		},
	}

	fn := func(t groupByMockStruct) string {
		return t.Name
	}
	for _, table := range tables {
		r := gfuncs.GroupBy(table.in, fn)
		if !reflect.DeepEqual(table.exp, r) {
			t.Errorf("Error on input %v: expected %v, got %v", table.in, table.exp, r)
		}
	}

}
