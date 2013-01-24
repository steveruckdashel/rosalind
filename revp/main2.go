package main

import (
	"fmt"
	"runtime"
)

const teststr = "TCAATGCATGCGGGTCTATATGCAT"
// 4 6
// 5 4
// 6 6
// 7 4
// 17 4
// 18 4
// 20 6
// 21 4
const realdata = "GCGCAATTTAAACGCATGATTGTATATCCCCATGGGAGGAATTTCTTTGTATTGACACATGGTTAACTTCAGATGAGTAGTTAGACGAACCATCTTGCTCCGCGCGATCTCGGATTACAGTCCCTTTTCCGTAAACCTACAACACATTGTCTCCAGGGTCAGGAAGCTTTATCTTTGTCGGCCGCGCACGTGATTAAAAATGAAATAGTATCGACTAGGCAGTGCGATTGGCCCTTTTGGGGGCTAGCTAGTGTGTGTAGTGGTTCCACCTATTGGAAAGTGCCACTGCGAGTGCGTCCGCCTCATCAAGCTTCCGTATGTGGGGTGCCTTGACAACTAGGTTACTGGCGTGAGTTCGGCGGCTGACCCATTCACAAATGCCAACCGTGAGGACGTGTATAAGTAACAGTAAGGTGACACAGTGGCTGGATCAGTACTTCGGGTACCTTAGGTGCCTGCGCGCGCGCATAGGATACCGTCCCGCTTATGACTCTAAGAGATACTGAACTTCGTGCATGAAAACACTTAGTTATATTCTGGAAACTCGGCGGTACCACTGCTGACTACGACGCCATTGGTGGCTGATTTACCTAATGACGGCGTGGGTATTGGTCTGTTGAAGTAAGGTCCGCATACTCCAAAAAGGCTTTCGGACTGCTCATGAACGTAACTCAATGACTTGGCTGCGAAATCATATTTGCTGCTCCAACCGCCACAATAAGGCCCCCTAATGGTTGATCGAATAGCACGTGATTAGCGTTAACGAGTGGATAACGCCGTTAGCTCCTTTCAAGCCGGGTGTTAATCGGCTGTACTCGTTCTGAGAAAGAACCCGGACTACTGGTGCACGCCTACCTAGGAAAGTGTTAGTCTTGTCCGTTTGTGATTACCCTCCGGCCACTTCCGTACCGGGCCCCCATATAT"


type SubStr struct {
	Len		int
	Index	int
	Str		string
	RP		chan bool
}

func (s *SubStr) String() (string) {
	return fmt.Sprintf("%d %d",s.Index, s.Len)
}

func SubStrGen(str string, min, max int) (<-chan SubStr) {
	ch := make(chan SubStr, 64)
	
	go func(){
		for wind:= min; wind<=max; wind++ {
			for i:=0; i+wind<= len(str); i++ {
				var ss SubStr
				ss.Str = str[i:i+wind]
				ss.Len = wind
				ss.Index = i+1
				ss.RP = IsRevPal(ss.Str)
				ch<- ss
			}
		}
		close(ch)
	}()
	
	return ch
}

func IsRevPal(s string) (chan bool) {
	ch := make(chan bool)
	
	go func(){
		for k,v := range s {
				rev := 'X'
			switch s[len(s) - k-1] {
				case 'G':
					rev = 'C'
				case 'C':
					rev = 'G'
				case 'T':
					rev = 'A'
				case 'A':
					rev = 'T'
			}
			if v != rev {
				ch<- false
				break
			}
		}
		ch<- true
	}()
	return ch
}


func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
//	ch:= SubStrGen(teststr,4,12)
	ch:= SubStrGen(realdata,4,12)
	for {
		s, ok := <-ch
		if !ok {break}
		if <-s.RP == true {
			fmt.Println(s.String())
		}
	}
}