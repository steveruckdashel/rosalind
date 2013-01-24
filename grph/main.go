package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"errors"
//	"log"
	"strings"
)

type FASTA struct {
	Name, DNA string
}

const O = 3

func BuildFASTA(name string) (chan rune, chan FASTA) {
	r := make(chan rune, 16)
	fchan := make(chan FASTA,1)
	fasta := FASTA{name,""}
	
	go func(){
		for v := range r {
			if v=='X' { break }
			fasta.DNA += string(v)
		}
		fchan<- fasta
	}()
	
	return r, fchan
}

func FASTAgen(fin, fout *os.File) {
	bin := bufio.NewReaderSize(fin, 4*1024)
	FASTAs := make(map[string]FASTA)
	var tempFasta FASTA
	
	var r rune
	err := errors.New("")
	for err!=io.EOF {
		for r!='>' {
			r, _, err = bin.ReadRune()
			if err!=nil { break }
		}
		if err==io.EOF { break }
		FASTAName,err := bin.ReadString('\n')
		FASTAName = strings.TrimSpace(FASTAName)
		if err==io.EOF { break }
		
		build, fasta := BuildFASTA(FASTAName)
		r= '.'
		for r!='>' {
			r, _, err = bin.ReadRune()
			if err==io.EOF { break }
			if strings.ContainsRune("ATGC",r) {build<- r}
		}
		build<- 'X'
		tempFasta = <-fasta
		FASTAs[tempFasta.Name]= tempFasta
	}
	
	for k,v := range FASTAs {
		l:=len(v.DNA)
		for k2, v2 := range FASTAs {
			if k!=k2 {
//				l2:=len(v2.DNA)
//				fmt.Println(l2)
				if v2.DNA[0:O]==v.DNA[l-O:l] {
					fmt.Fprintf(fout, "%s %s\n",k,k2)
				}
			}
		}
	}
}

func main() {
	FASTAgen(os.Stdin,os.Stdout)
}