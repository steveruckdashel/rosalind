package main

import (
	"fmt"
	"os"
	"log"
	"io"
	"errors"
)

func Transcribe(fin, fout *os.File) {
	char := make([]byte,1)
	err:= errors.New("")
	n:=0
	for err!=io.EOF {
		n,err =fin.Read(char)
		if n>0 {
			switch char[0] {
				case 'T': fmt.Fprintf(fout,"%c",'U')
				default: fmt.Fprintf(fout,"%c",char[0])
			}
		}
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