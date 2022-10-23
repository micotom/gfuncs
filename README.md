# GFuncs

Implementation of slice functions for Go 1.18+ inspired by the [Kotlin Collection extensions](https://kotlinlang.org/docs/collection-operations.html).

## Download

```
go get github.com/micotom/gfuncs
```

## (Some of the) Supported Functions

`Fold` : `[]T, V, (V, T -> V) -> V`
```go 
func Fold[T any, V any](slice []T, initial V, fn func(acc V, t T) V) V
```

`FoldIndexed` : `[]T, V, (int, T, V -> V) -> V`
```go
func FoldIndexed[T any, V any](slice []T, initial V, fn func(index int, acc V, t T) V) V
```

`SortBy` : `[]T, (T -> V) -> []T`
```go
func SortBy[T any, V constraints.Ordered](slice []T, fn func(t T) V) []T
```

`Filter` : `[]T, (T -> bool) -> []T`
```go
func Filter[T any](slice []T, fn func(t T) bool) []T
```

`FilterIndexed` : `[]T, (int, T -> bool) -> []T`
```go
func FilterIndexed[T any](slice []T, fn func(i int, t T) bool) []T
```

`GroupBy` : `[]T, (T -> V) -> map[V][]T`
```go
func GroupBy[T any, V comparable](slice []T, fn func(t T) V) map[V][]T
```

`SumBy` : `[]T, (T -> V) -> *V, error`
```go
func SumBy[T any, V Number](slice []T, fn func(t T) V) (*V, error)
```

`Map` : `[]T, (T -> V) -> []V`
```go
func Map[T any, V any](slice []T, fn func(t T) V) []V
```

`MapIndexed` : `[]T, (int, T -> V) -> []V`
```go
func MapIndexed[T any, V any](slice []T, fn func(i int, t T) V) []V 
```

`Flatten` : `[][]T -> []T`
```go
func Flatten[T any](slice [][]T) []T
```

`Zip` : `[]T, []V -> []Tuple[T, V]`
```go
func Zip[T any, V any](sliceT []T, sliceV []V) []Tuple[T, V]
```

`Find` : `[]T, (T -> bool) -> *T`
```go
func Find[T any](slice []T, fn func(t T) bool) *T
```

`Any` : `[]T, (T -> bool) -> bool`
```go
func Any[T any](slice []T, fn func(t T) bool) bool
```

`All` : `[]T, (T -> bool) -> bool`
```go
func All[T any](slice []T, fn func(t T) bool) bool
```