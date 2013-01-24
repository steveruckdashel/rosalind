package main

import (
	"fmt"
	"math"
	"runtime"
)

// const s = "ACGATACAA"
const s = "CGGCGGGCTTCCGTCCACCAGGTATGCACATCGCTGACTTTATACAAGTCAAGGCTGGGCTTAGGTCACGAGACTCGCATTTTCCT"
// const A = {0.129,0.287,0.423,0.476,0.641,0.742,0.783}

func Logcumulator(in <-chan float64) (<-chan float64) {
	out := make(chan float64)
	
	go func(){
		acc := 0.0
		for {
			x, ok := <-in
			if !ok {break}
			acc += math.Log10(x)
		}
		out<- acc
	}()
	
	return out
}

func checkProb(str string, p float64) (float64) {
	probs := make(chan float64, 16)
	ch := Logcumulator(probs)
	
	for _,v := range str {
		switch v {
			case 'G': fallthrough
			case 'C': probs<- p / 2.0
			case 'T': fallthrough
			case 'A': probs<- (1.0 - p) / 2.0
		}
	}
	
	close(probs)
	return <-ch
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	
	A := []float64{0.093, 0.136, 0.183, 0.249, 0.289, 0.377, 0.436, 0.469, 0.524, 0.577, 0.643, 0.685, 0.735, 0.805, 0.876, 0.902}
	// -76.023 -69.088 -63.977 -59.117 -57.003 -53.816 -52.533 -52.065 -51.653 -51.686 -52.349 -53.177 -54.666 -58.005 -63.947 -67.336
	B := make([]float64, len(A))
	for i,a := range A {
		B[i] = checkProb(s,a)
	}
	
	for _,b := range B {
		fmt.Printf("%.3f ", b)
	}
	fmt.Print("\n")
}
