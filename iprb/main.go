package main

import (
	"fmt"
//	"math/rand"
//	"time"
)

/*
func RandGenerator() (<-chan float64) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ch := make(chan float64, 128)
	go func(){
		for i:=0; i<10000; i++ {
			ch<- r.Float64()
		}
		close(ch)
	}()
	return ch
}
*/

func first(k,m,n int) ([](chan float64)) {
	ch := make([]chan float64,9)
	for i := range ch {ch[i] = make(chan float64)}
	
	go func(){
		for i,c:= range ch {
			prob := 0.0
			switch i {
				case 0:
					prob = float64(k)/ float64(k+m+n) 
					prob *= float64(k-1)/ float64(k+m+n-1)
				case 1:
					prob = float64(k)/ float64(k+m+n)
					prob *= float64(m)/ float64(k+m+n-1)
				case 2:
					prob = float64(k)/ float64(k+m+n)
					prob *= float64(n)/ float64(k+m+n-1)
				case 3:
					prob = float64(m)/ float64(k+m+n)
					prob *= float64(k)/ float64(k+m+n-1)
				case 4:
					prob = float64(m)/ float64(k+m+n)
					prob *= float64(m-1)/ float64(k+m+n-1)
				case 5:
					prob = float64(m)/ float64(k+m+n)
					prob *= float64(n)/ float64(k+m+n-1)
				case 6:
					prob = float64(n)/ float64(k+m+n)
					prob *= float64(k)/ float64(k+m+n-1)
				case 7:
					prob = float64(n)/ float64(k+m+n)
					prob *= float64(m)/ float64(k+m+n-1)
				case 8:
					prob = float64(n)/ float64(k+m+n)
					prob *= float64(n-1)/ float64(k+m+n-1)

			}
			c<- prob
		}
		
	}()
	
	return ch
}


func main() {
	probs := []float64{1.0,1.0,1.0,
					   1.0,.75,0.5,
					   1.0,0.5,0.0}
	
	// ans := first(2,2,2)
	// expect := 0.78333
	ans := first(22,26,23)
	// expect := 0.74517

	
	for k,v := range ans {
		sum += <-v * probs[k]
	}
	fmt.Println(sum)
}