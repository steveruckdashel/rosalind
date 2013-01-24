package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strings"
)



func Hamming(fin *os.File) (dist int) {
	dist=0
	bin := bufio.NewReaderSize(fin, 1024)
	one,_ := bin.ReadString('\n')
	two,_ := bin.ReadString('\n')
	
	one = strings.TrimSpace(one)
	two = strings.TrimSpace(two)
	
//	fmt.Println(one)
//	fmt.Println(two)
	
	for k,v := range two {
		if byte(v)!=one[k] {
			dist++
		}
	}
	
	return dist
}

func main() {
	if (len(os.Args)<2) {
		log.Fatal("Need 1 file names (input & output)")
	}
	fin, err := os.Open(os.Args[1])
	if err!=nil {
		log.Fatal(err)
	}
	defer fin.Close()
	
	fmt.Println(Hamming(fin))
}