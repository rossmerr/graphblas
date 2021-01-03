# GraphBLAS

![Go](https://github.com/rossmerr/graphblas/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/rossmerr/graphblas)](https://goreportcard.com/report/github.com/rossmerr/graphblas)
[![Read the Docs](https://pkg.go.dev/badge/golang.org/x/pkgsite)](https://pkg.go.dev/github.com/rossmerr/graphblas)

A sparse linear algebra library implementing may of the ideas from the [GraphBLAS Forum](https://graphblas.github.io/) in Go.

Supports float32 and float64 types

```go
array := [][]float64{
		[]float64{0, 0, 0, 1, 0, 0, 0},
		[]float64{1, 0, 0, 0, 0, 0, 0},
		[]float64{0, 0, 0, 1, 0, 1, 1},
		[]float64{1, 0, 0, 0, 0, 0, 1},
		[]float64{0, 1, 0, 0, 0, 0, 1},
		[]float64{0, 0, 1, 0, 1, 0, 0},
		[]float64{0, 1, 0, 0, 0, 0, 0},
    }
    
g := doubleprecision.NewDenseMatrixFromArray(array)

atx := breadthfirst.Search(context.Background(), g, 3, func(i doubleprecision.Vector) bool {
    return i.AtVec(5) == 1
})
```    