package sort

import (
	"context"

	"github.com/rossmerr/graphblas"
)

func BubbleRow(ctx context.Context, a graphblas.MatrixRune) graphblas.MatrixRune {
	size := a.Rows()
	result := a.CopyLogical()
	for j := 0; j < size-1; j++ {
		for i := j + 1; i < size; i++ {
			vj := result.RowsAt(j)
			vi := result.RowsAt(i)
			if graphblas.Compare(ctx, vj, vi) > 0 {
				enumerator := vi.Enumerate()
				for enumerator.HasNext() {
					c, _, v := enumerator.Next()
					result.Set(j, c, v)
				}

				enumerator = vj.Enumerate()
				for enumerator.HasNext() {
					c, _, v := enumerator.Next()
					result.Set(i, c, v)
				}
			}
		}
	}

	return result
}

func BubbleColumns(ctx context.Context, a graphblas.MatrixRune) graphblas.MatrixRune {
	size := a.Columns()
	result := a.CopyLogical()
	for j := 0; j < size-1; j++ {
		for i := j + 1; i < size; i++ {
			vj := result.ColumnsAt(j)
			vi := result.ColumnsAt(i)
			if graphblas.Compare(ctx, vj, vi) > 0 {
				enumerator := vi.Enumerate()
				for enumerator.HasNext() {
					c, _, v := enumerator.Next()
					result.Set(c, j, v)
				}

				enumerator = vj.Enumerate()
				for enumerator.HasNext() {
					c, _, v := enumerator.Next()
					result.Set(c, i, v)
				}
			}
		}
	}

	return result
}
