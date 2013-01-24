package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"bufio"
)



func Transcribe(fin, fout *os.File) {
	bin := bufio.NewReaderSize(fin, 16 * 1024)
	forward,_ := bin.ReadString('.')
	
	forward = strings.Replace(forward, "A","1",-1)
	forward = strings.Replace(forward, "T","2",-1)
	forward = strings.Replace(forward, "C","3",-1)
	forward = strings.Replace(forward, "G","4",-1)
	
	forward = strings.Replace(forward, "1","T",-1)
	forward = strings.Replace(forward, "2","A",-1)
	forward = strings.Replace(forward, "3","G",-1)
	forward = strings.Replace(forward, "4","C",-1)
	
	for i:=len(forward)-1; i>=0; i-- {
		fmt.Fprintf(fout,"%c",forward[i])
	}
}

func main() {
	if (len(os.Args)<3) {
		log.Fatal("Need 2 file names (input & output)")
	}
	fin, err := os.Open(os.Args[1])
	if err!=nil {
		log.Fatal(err)
	}
	defer fin.Close()
	
	fout, err := os.Create(os.Args[2])
	if err!=nil {
		log.Fatal(err)
	}
	defer fout.Close()
	
	Transcribe(fin,fout)
}