package main

import (
	"fmt"
)

// 1 AA-AA -> 1.0
// 2 AA-Aa -> 1.0
// 3 AA-aa -> 1.0
// 4 Aa-Aa -> .75
// 5 Aa-aa -> 0.5
// 6 aa-aa -> 0.0
//
// 1 0 0 1 0 1
// 3.5

func main() {
	probs := []float64{1.0,1.0,1.0,.75,0.5,0.0}
//	pop := []float64{1, 0, 0, 1, 0, 1}
	pop := []float64{19834, 16929, 17870, 18250, 17217, 16453}
	
	avg := 0.0
	
	for k,v := range pop {
		avg+= v * probs[k]
	}
	fmt.Println(avg*2)
}
