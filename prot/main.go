package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"errors"
	"log"
)

func Transcode(fin, fout *os.File) {
	codons := map[string]string{"UUU":"F",		"CUU":"L",	"AUU":"I",	"GUU":"V",
								"UUC":"F",		"CUC":"L",	"AUC":"I",	"GUC":"V",
								"UUA":"L",		"CUA":"L",	"AUA":"I",	"GUA":"V",
								"UUG":"L",		"CUG":"L",	"AUG":"M",	"GUG":"V",
								"UCU":"S",		"CCU":"P",	"ACU":"T",	"GCU":"A",
								"UCC":"S",		"CCC":"P",	"ACC":"T",	"GCC":"A",
								"UCA":"S",		"CCA":"P",	"ACA":"T",	"GCA":"A",
								"UCG":"S",		"CCG":"P",	"ACG":"T",	"GCG":"A",
								"UAU":"Y",		"CAU":"H",	"AAU":"N",	"GAU":"D",
								"UAC":"Y",		"CAC":"H",	"AAC":"N",	"GAC":"D",
								"UAA":"Stop",	"CAA":"Q",	"AAA":"K",	"GAA":"E",
								"UAG":"Stop",	"CAG":"Q",	"AAG":"K",	"GAG":"E",
								"UGU":"C",		"CGU":"R",	"AGU":"S",	"GGU":"G",
								"UGC":"C",		"CGC":"R",	"AGC":"S",	"GGC":"G",
								"UGA":"Stop",	"CGA":"R",	"AGA":"R",	"GGA":"G",
								"UGG":"W",		"CGG":"R",	"AGG":"R",	"GGG":"G",}

	bin := bufio.NewReaderSize(fin, 4*1024)
	var triple string
	err:= errors.New("")
	n:=0
	var r rune
	for err!=io.EOF {
		triple = ""
		for i:=0; i<3; i+=n {
			r, n, err = bin.ReadRune()
			if err!=nil {
				break
			}
			triple += string(r)
		}
		
		code := codons[triple]
		if code=="Stop" {
			break
		}
		fmt.Fprint(fout,code)
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
	
	Transcode(fin,fout)
}