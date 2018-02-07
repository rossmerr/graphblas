// Copyright (c) 2018 Ross Merrigan
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"runtime/pprof"
	"testing"

	GraphBLAS "github.com/RossMerr/Caudex.GraphBLAS"
)

var cpuprofile = flag.String("cpuprofile", "cpu.prof", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	for _, bm := range benchmarks {
		seconds, err := runBenchmarkFor(bm.fn)
		if err != nil {
			log.Fatalf("%s %s", bm.name, err)
		}
		print_perf(bm.name, seconds)
	}
}

func print_perf(name string, time float64) {
	fmt.Printf("go,%v,%v\n", name, time*1000)
}

func runBenchmarkFor(fn func(*testing.B)) (seconds float64, err error) {
	bm := testing.Benchmark(fn)
	if (bm == testing.BenchmarkResult{}) {
		return 0, errors.New("failed")
	}
	return bm.T.Seconds() / float64(bm.N), nil
}

func BenchmarkMatrix(b *testing.B) {
	for _, fn := range benchmarks {
		fn.fn(b)
	}
}

var benchmarks = []struct {
	name string
	fn   func(*testing.B)
}{

	{
		name: "iteration_pi_sum",
		fn: func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				if math.Abs(pisum()-1.644834071848065) >= 1e-6 {
					b.Fatal("pi_sum out of range")
				}
			}
		},
	},
	{
		name: "matrix_multiply",
		fn: func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				c := randmatmul(1)
				v := c.At(0, 0)
				if !(v >= 0) {
					b.Fatal("assert failed")
				}
			}
		},
	},
}

func pisum() float64 {
	var sum float64
	for i := 0; i < 500; i++ {
		sum = 0.0
		for k := 1.0; k <= 10000; k += 1 {
			sum += 1.0 / (k * k)
		}
	}
	return sum
}

func randmatmul(n int) GraphBLAS.Matrix {
	aData := make([][]float64, n)
	for r := range aData {
		aData[r] = make([]float64, n)
		for c := range aData {
			aData[r][c] = rand.Float64()

		}
	}
	a := GraphBLAS.NewDenseMatrixFromArray(aData)

	bData := make([][]float64, n)
	for r := range bData {
		bData[r] = make([]float64, n)
		for c := range bData {
			bData[r][c] = rand.Float64()

		}
	}
	b := GraphBLAS.NewDenseMatrixFromArray(bData)

	c := a.Multiply(b)
	return c
}
