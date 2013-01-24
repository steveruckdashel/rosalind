package main

import (
	"fmt"
	"strings"
	"os"
	"log"
	"io"
	"bufio"
)


func GenSubStrings(str string) (chan string) {
	s:= make(chan string, 16)
	old := make(map[string]bool)
	
	go func(){
		for l:=len(str); l>=2; l-- {
			for i:=0; i+l <= len(str); i++ {
				sub := str[i:i+l]
				if old[sub]!= true {
					old[sub] = true
					s<- sub
				}
			}
		}
		s<- "X"
	}()
	
	return s
}

func SubStringSieve(str string, subs chan string) (chan string) {
	s:= make(chan string, 16)
	old := make(map[string]bool)
	
	go func(){
		var sub string = ""
		for sub!= "X" {
			sub= <-subs
			if strings.Contains(str, sub) {
				if old[sub]!= true {
					old[sub] = true
					s<- sub
				}
			}
		}
		s<- "X"
	}()
	
	return s
}



func main() {
	var ch chan string
	bin := bufio.NewReaderSize(os.Stdin, 16*1024)
	DNA,err := bin.ReadString('\n')
	DNA = strings.TrimSpace(DNA)
	if err!=nil {
		log.Fatal(err) 
	} else {ch=GenSubStrings(DNA)}
	
	for {
		DNA,err := bin.ReadString('\n')
		if err!=nil {
			if err==io.EOF {break}
			log.Fatal(err) 
		}
		DNA = strings.TrimSpace(DNA)
		ch=SubStringSieve(DNA,ch)
	}
	
	for s:= range ch {
		if s== "X" {break}
		fmt.Println(s)
	}
}
