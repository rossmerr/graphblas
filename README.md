# GraphBLAS

![Go](https://github.com/rossmerr/graphblas/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/rossmerr/graphblas)](https://goreportcard.com/report/github.com/rossmerr/graphblas)
[![Read the Docs](https://pkg.go.dev/badge/golang.org/x/pkgsite)](https://pkg.go.dev/github.com/rossmerr/graphblas)

A sparse linear algebra library implementing may of the ideas from the [GraphBLAS Forum](https://graphblas.github.io/) in Go.

Sparse Matrix Formats:
Compressed Sparse Row (CSR)
Compressed Sparse Column (CSC)
Sparse Vector

Supports bool | int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64

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

g := graphblas.NewDenseMatrixFromArray(array)

atx := breadthfirst.Search[float64](context.Background(), g, 3, func(i graphblas.Vector[float64]) bool {
    return i.AtVec(5) == 1
})
```
