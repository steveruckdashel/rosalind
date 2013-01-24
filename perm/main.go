package main

import (
	"fmt"
	"strings"
)

func fact(f int) (int) {
	result := 1
	for f>0 {
		result *=f
		f--
	}
	return result
}

func FanOut(car, cdr string) {
	if cdr=="" {
		
		fmt.Println(strings.Join(strings.Split(car,""), " "))
		return
	}
	for k,v:= range cdr {
		end := cdr[0:k] + cdr[k+1:len(cdr)]
		FanOut(car+string(v),end)
	}
	
}

func main() {
	num := 7
	numOfPerms := fact(num)
	fmt.Println(numOfPerms)
	
	FanOut("","1234567"[0:num])
}