package operators_test

import (
	"testing"

	"github.com/rossmerr/graphblas/operators"
)

func Test_UnaryOp(t *testing.T) {
	operators.LogicalInverse.UnaryOp()
}

func Test_Operator(t *testing.T) {
	operators.LOR.Operator()
	operators.LogicalInverse.Operator()
	operators.FirstArgument[float32]().Operator()

}

func Test_BinaryOp(t *testing.T) {
	operators.LOR.BinaryOp()
	operators.FirstArgument[float32]().BinaryOp()

}

func Test_Semigroup(t *testing.T) {
	operators.LOR.Semigroup()
	operators.FirstArgument[float32]().Semigroup()

}
