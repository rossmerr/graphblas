package sort

import (
	"context"

	"github.com/rossmerr/graphblas"
)

func Bubble(ctx context.Context, a graphblas.MatrixRune, n int) graphblas.MatrixRune {
	result := a.CopyLogical()
	for j := 0; j < n-1; j++ {
		for i := j + 1; i < n; i++ {
			aj := a.RowsAt(j)
			vi := a.RowsAt(i)
			if graphblas.Compare(ctx, aj, vi) > 0 {
				enumerator := vi.Enumerate()
				if enumerator.HasNext() {
					_, c, v := enumerator.Next()
					result.Set(j, c, v)
				}

				enumerator = aj.Enumerate()
				if enumerator.HasNext() {
					_, c, v := enumerator.Next()
					result.Set(i, c, v)
				}
			}
		}
	}

	return result
}
