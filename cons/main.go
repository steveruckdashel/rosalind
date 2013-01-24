package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"bufio"
)

type KVP struct {
	K	int
	V	rune
}

var lock chan int

func Mapper(strlen int) (chan KVP) {
	c := make(chan KVP, 4)
	m := make(map[rune]([]int))
	
	m['A'] = make([]int,strlen)
	m['C'] = make([]int,strlen)
	m['G'] = make([]int,strlen)
	m['T'] = make([]int,strlen)
	max := make([]KVP,strlen)
	
	go func(){
		for pair := range c {
			if pair.V == 'X' {
				break
			}
			m[pair.V][pair.K]++
		}
		
		
		
		for _,v := range []rune{'A','C','G','T'} {
			fmt.Print(string(v)+ ":")
			for k,n := range m[v] {
				
				if n>max[k].K {
					max[k]=KVP{n,v}
				}
				fmt.Printf(" %d" ,n)
			}
			fmt.Println()
		}
		
		for _,v := range max {
			fmt.Printf("%c",v.V)
		}
		
//		fmt.Println("A:",m['A'])
//		fmt.Println("C:",m['C'])
//		fmt.Println("G:",m['G'])
//		fmt.Println("T:",m['T'])
		lock<- 1
	}()
	
	
	return c
} 

func Transcribe(fin *os.File) {
	lock = make(chan int)
	bin := bufio.NewReaderSize(fin, 16 * 1024)
	
	header,err := bin.ReadString('\n')
	if err!=nil {
		log.Fatal(err)
	}
	header = strings.TrimSpace(header)
	count := Mapper(len(header))
	for k,v := range header {
		count<- KVP{k,v}
	}
	
	for ;; {
		str,err := bin.ReadString('\n')
		str = strings.TrimSpace(str)
		for k,v := range str {
			count<- KVP{k,v}
		}
		if err!=nil {
			count<- KVP{0,'X'}
			break
		}
	}
}

func main() {
	if (len(os.Args)<2) {
		log.Fatal("Need 1 file names (input)")
	}
	fin, err := os.Open(os.Args[1])
	if err!=nil {
		log.Fatal(err)
	}
	defer fin.Close()
	
	Transcribe(fin)
	
	<-lock
}
