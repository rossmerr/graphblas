// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package f32

import (
	"log"
	"reflect"
)

var sparseMatrixRegistry = make(map[string]reflect.Type)

// RegisterMatrix add's the sparse matrix to the registry
func RegisterMatrix(matrix reflect.Type) {
	if _, found := sparseMatrixRegistry[matrix.Name()]; found {
		log.Panicf("Already registered Matrix %q.", matrix.Name())
	}
	sparseMatrixRegistry[matrix.Name()] = matrix

}

// IsSparseMatrix is 's' a sparse matrix
func IsSparseMatrix(s Matrix) bool {
	t := reflect.TypeOf(s).Elem()
	_, found := sparseMatrixRegistry[t.Name()]
	return found
}
