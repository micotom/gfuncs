package gfuncs

import (
	"errors"

	"golang.org/x/exp/constraints"
)

// Missing: associateWith, associateBy, findLast

type Number interface {
	constraints.Integer | constraints.Float
}

// Folds the slice from the left to a single value by applying the function argument with the initial value as starting point.
func Fold[T any, V any](slice []T, initial V, fn func(acc V, t T) V) V {
	for _, t := range slice {
		initial = fn(initial, t)
	}
	return initial
}

// Folds the slice from the left to a single value by applying the function argument with the initial value as starting while also handing over the current index of the slice the function is called upon.
func FoldIndexed[T any, V any](slice []T, initial V, fn func(index int, acc V, t T) V) V {
	for i, t := range slice {
		initial = fn(i, initial, t)
	}
	return initial
}

// Sorts the slice by determine the order based on the passede function for that specific element of the slice. Sorting algorithm is quicksort.
func SortBy[T any, V constraints.Ordered](slice []T, fn func(t T) V) []T {

	var sort func([]T, int, int)
	sort = func(arr []T, start int, end int) {
		if (end - start) < 1 {
			return
		}

		pivot := arr[end]
		splitIndex := start
		for i := start; i < end; i++ {
			if fn(arr[i]) < (fn(pivot)) {
				temp := arr[splitIndex]
				arr[splitIndex] = arr[i]
				arr[i] = temp
				splitIndex++
			}
		}
		arr[end] = arr[splitIndex]
		arr[splitIndex] = pivot

		sort(arr, start, splitIndex-1)
		sort(arr, splitIndex+1, end)
	}

	r := make([]T, len(slice))

	for i, t := range slice {
		r[i] = t
	}

	sort(r, 0, len(r)-1)

	return r
}

// Returns a new slice with all elements of the base slice satisfying the applied function to return true.
func Filter[T any](slice []T, fn func(t T) bool) []T {
	foldFn := func(acc []T, t T) []T {
		if fn(t) {
			return append(acc, t)
		} else {
			return acc
		}
	}
	return Fold(slice, []T{}, foldFn)
}

// Returns a new slice with all elements of the base slice satisfying the applied function to return true. The function will retrieve the element and its index in the base slice.
func FilterIndexed[T any](slice []T, fn func(i int, t T) bool) []T {
	foldFn := func(i int, acc []T, t T) []T {
		if fn(i, t) {
			return append(acc, t)
		} else {
			return acc
		}
	}
	return FoldIndexed(slice, []T{}, foldFn)
}

// Returns a map with the keys derived from the function passed
func GroupBy[T any, V comparable](slice []T, fn func(t T) V) map[V][]T {
	foldFn := func(acc map[V][]T, t T) map[V][]T {
		k := fn(t)
		if val, kPresent := acc[k]; kPresent {
			acc[k] = append(val, t)
		} else {
			acc[k] = []T{t}
		}
		return acc
	}
	return Fold(slice, make(map[V][]T), foldFn)
}

func SumBy[T any, V Number](slice []T, fn func(t T) V) (*V, error) {
	if len(slice) == 0 {
		return nil, errors.New("Cannot sum slice of length 0")
	}
	var initial V
	foldFn := func(acc V, t T) V {
		return acc + fn(t)
	}
	r := Fold(slice, initial, foldFn)
	return &r, nil
}

func Map[T any, V any](slice []T, fn func(t T) V) []V {
	foldFn := func(acc []V, t T) []V {
		return append(acc, fn(t))
	}
	return Fold(slice, make([]V, 0, len(slice)), foldFn)
}

func MapIndexed[T any, V any](slice []T, fn func(i int, t T) V) []V {
	foldFn := func(i int, acc []V, t T) []V {
		return append(acc, fn(i, t))
	}
	return FoldIndexed(slice, make([]V, 0, len(slice)), foldFn)
}

func Flatten[T any](slice [][]T) []T {
	foldFn := func(acc []T, t []T) []T {
		for _, t := range t {
			acc = append(acc, t)
		}
		return acc
	}
	return Fold(slice, []T{}, foldFn)
}

type Tuple[T any, V any] struct {
	First  T
	Second V
}

func Zip[T any, V any](sliceT []T, sliceV []V) []Tuple[T, V] {
	if len(sliceT) < len(sliceV) {
		r := make([]Tuple[T, V], len(sliceT))
		return FoldIndexed(sliceT, r, func(index int, acc []Tuple[T, V], t T) []Tuple[T, V] {
			acc[index] = Tuple[T, V]{t, sliceV[index]}
			return acc
		})
	} else {
		r := make([]Tuple[T, V], len(sliceV))
		return FoldIndexed(sliceV, r, func(index int, acc []Tuple[T, V], v V) []Tuple[T, V] {
			acc[index] = Tuple[T, V]{sliceT[index], v}
			return acc
		})
	}
}

func Find[T any](slice []T, fn func(t T) bool) *T {
	for _, t := range slice {
		if fn(t) {
			return &t
		}
	}
	return nil
}

func Any[T any](slice []T, fn func(t T) bool) bool {
	return Find(slice, fn) != nil
}

func All[T any](slice []T, fn func(t T) bool) bool {
	for _, t := range slice {
		if !fn(t) {
			return false
		}
	}
	return true
}

func None[T any](slice []T, fn func(t T) bool) bool {
	return !Any(slice, fn)
}

func AsSet[T any, V comparable](slice []T, fn func(t T) V) []T {
	m := make(map[V]T)
	for _, t := range slice {
		m[fn(t)] = t
	}
	r := make([]T, 0, len(slice))
	for _, t := range m {
		r = append(r, t)
	}
	return r
}

func Partition[T any](slice []T, fn func(t T) bool) [][]T {
	left := make([]T, 0, len(slice))
	right := make([]T, 0, len(slice))
	foldFn := func(acc [][]T, t T) [][]T {
		if fn(t) {
			acc[0] = append(acc[0], t)
		} else {
			acc[1] = append(acc[1], t)
		}
		return acc
	}
	return Fold(slice, [][]T{left, right}, foldFn)
}

func Reverse[T any](slice []T) []T {
	rev := make([]T, len(slice))
	for i, t := range slice {
		rev[len(slice)-1-i] = t
	}
	return rev
}

func IndexOf[T any](slice []T, t T, fn func(t T) bool) int {
	for i, ts := range slice {
		if fn(t) == fn(ts) {
			return i
		}
	}
	return -1
}
