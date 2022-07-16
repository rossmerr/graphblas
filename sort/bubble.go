package sort

import (
	"context"

	"github.com/rossmerr/graphblas"
)

func BubbleColumns(ctx context.Context, a graphblas.MatrixRune, n int) graphblas.MatrixRune {
	result := a.CopyLogical()
	for j := 0; j < n-1; j++ {
		for i := j + 1; i < n; i++ {
			aj := result.ColumnsAt(j)
			vi := result.ColumnsAt(i)
			if graphblas.Compare(ctx, aj, vi) > 0 {
				enumerator := vi.Enumerate()
				for enumerator.HasNext() {
					_, c, v := enumerator.Next()
					result.Set(j, c, v)
				}

				enumerator = aj.Enumerate()
				for enumerator.HasNext() {
					_, c, v := enumerator.Next()
					result.Set(i, c, v)
				}
			}
		}
	}

	return result
}

func BubbleRow(ctx context.Context, a graphblas.MatrixRune, n int) graphblas.MatrixRune {
	result := a.CopyLogical()
	for j := 0; j < n-1; j++ {
		for i := j + 1; i < n; i++ {
			aj := result.RowsAt(j)
			vi := result.RowsAt(i)
			if graphblas.Compare(ctx, aj, vi) > 0 {
				enumerator := vi.Enumerate()
				for enumerator.HasNext() {
					_, c, v := enumerator.Next()
					result.Set(j, c, v)
				}

				enumerator = aj.Enumerate()
				for enumerator.HasNext() {
					_, c, v := enumerator.Next()
					result.Set(i, c, v)
				}
			}
		}
	}

	return result
}
