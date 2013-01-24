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

func CountDNA() (chan rune, chan map[rune]int) {
	r := make(chan rune, 16)
	tot := make(chan map[rune]int,1)
	
	t := map[rune]int{'A':0,'C':0,'G':0,'T':0}
	
	go func(){
		for v := range r {
			if v=='X' { break }
			t[v]++
		}
		tot<- t
	}()
	
	return r, tot
}

func FASTAgen(fin, fout *os.File) {
	bin := bufio.NewReaderSize(fin, 4*1024)
	FASTAs := make(map[string]map[rune]int)

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
		
		count, total := CountDNA()
		r= '.'
		for r!='>' {
			r, _, err = bin.ReadRune()
			if err==io.EOF { break }

			switch r {
				case 'A': fallthrough
				case 'T': fallthrough
				case 'G': fallthrough
				case 'C': count<- r
			}

		}
		count<- 'X'
		FASTAs[FASTAName]= <-total
	}
	max := 0.0
	maxname := ""
	for k,v := range FASTAs {
		gc := v['G'] + v['C']
		sum := gc + v['A'] + v['T']
		pcent := float64(gc)/float64(sum)
		if pcent > max {
			max = pcent
			maxname = k
		}
	}
	fmt.Fprintf(fout,"%s\n%.2f%%",maxname,max*100)
}

func main() {
	FASTAgen(os.Stdin,os.Stdout)
}