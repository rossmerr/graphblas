// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package predecessor

import GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS"

// function x = predecessor(G, v)
// % Predecessors of a node in a graph

// x = sparse(length(G), 1);
// xold = x;
// x(v) = 1; % Start BFS from v.

// while x ∼= xold
// 	xold = x;
// 	x = x | G * x;
// end

// Predecessors of a node in a graph
// v are nodes from which v is reachable and are found by breadth-first search in g
func Predecessors(g GraphBLAS.Matrix, v GraphBLAS.Vector) {
	var x GraphBLAS.Matrix = v
	var xold GraphBLAS.Matrix
	for x.NotEqual(xold) {
		xold = x
		x = g.Multiply(x)
	}
}
