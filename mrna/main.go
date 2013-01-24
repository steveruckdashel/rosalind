package main

import (
	"fmt"
	"bufio"
//	"log"
	"os"
	"strings"
	"math/big"
)
	
	

func main() {
	codons := map[string]([]string){
		"A":{"GCU","GCC","GCA","GCG"},
		"C":{"UGC","UGU"},
		"D":{"GAU","GAC"},
		"E":{"GAA","GAG"},
		"F":{"UUU","UUC"},
		"G":{"GGU","GGG","GGA","GGC"},
		"H":{"CAU","CAC"},
		"I":{"AUU","AUC","AUA"},
		"K":{"AAA","AAG"},
		"L":{"CUC","CUU","UUA","CUA","UUG","CUG"},
		"M":{"AUG"},
		"N":{"AAC","AAU"},
		"P":{"CCC","CCA","CCU","CCG"},
		"Q":{"CAA","CAG"},
		"R":{"CGG","CGU","AGG","CGC","CGA","AGA"},
		"S":{"UCU","UCC","AGU","AGC","UCA","UCG"},
		"T":{"ACU","ACC","ACA","ACG"},
		"V":{"GUC","GUU","GUA","GUG"},
		"W":{"UGG"},
		"Y":{"UAU","UAC"},
		"Stop":{"UAA","UAG","UGA"},}
	bin := bufio.NewReaderSize(os.Stdin, 2*1024)
	mRNA,_ := bin.ReadString('\n')
	mRNA = strings.TrimSpace(mRNA)
	
	possible:= big.NewInt(1)
	fmt.Printf("%d\n",possible)
	for _,v := range mRNA {
		count := big.NewInt(int64(len(codons[string(v)])))
		possible.Mul(possible,count)
		fmt.Printf("%v\n",possible)
	}
	possible.Mul(possible,big.NewInt(3))
	fmt.Println(possible)
}

