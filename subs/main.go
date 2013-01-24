package main

import (
	"fmt"
	"os"
	"log"
	"strings"
	"bufio"
)

func Transcribe(fin *os.File) {
	bin := bufio.NewReaderSize(fin, 16 * 1024)

	str,_ := bin.ReadString('\n')
	str = strings.TrimSpace(str)
	krnl,_ := bin.ReadString('\n')
	krnl = strings.TrimSpace(krnl)
	
	for i:=len(krnl); i<=len(str); i++ {
		if string(str[i-len(krnl):i])==krnl {
			fmt.Printf("%d ",i-len(krnl)+1)
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
}