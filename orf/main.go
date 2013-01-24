package main

import (
	"fmt"
	"strings"
)

var rna map[string]string = map[string]string{"UUU":"F","CUU":"L","AUU":"I","GUU":"V",
							"UUC":"F","CUC":"L","AUC":"I","GUC":"V",
							"UUA":"L","CUA":"L","AUA":"I","GUA":"V",
							"UUG":"L","CUG":"L","AUG":"M","GUG":"V",
							"UCU":"S","CCU":"P","ACU":"T","GCU":"A",
							"UCC":"S","CCC":"P","ACC":"T","GCC":"A",
							"UCA":"S","CCA":"P","ACA":"T","GCA":"A",
							"UCG":"S","CCG":"P","ACG":"T","GCG":"A",
							"UAU":"Y","CAU":"H","AAU":"N","GAU":"D",
							"UAC":"Y","CAC":"H","AAC":"N","GAC":"D",
							"UAA":"","CAA":"Q","AAA":"K","GAA":"E",
							"UAG":"","CAG":"Q","AAG":"K","GAG":"E",
							"UGU":"C","CGU":"R","AGU":"S","GGU":"G",
							"UGC":"C","CGC":"R","AGC":"S","GGC":"G",
							"UGA":"","CGA":"R","AGA":"R","GGA":"G",
							"UGG":"W","CGG":"R","AGG":"R","GGG":"G"}

func ParseORF(str string) (string) {
	s := ""
	for {
		next := rna[str[:3]]
		s += next
		if next == "" {break}
		str = str[3:]
		if len(str)<3 {return ""}
	}
	s += "\n"
	return s
}

func main() {
	s := "AGCCATGTAGCTAACTCAGGTTACATGGGGATGACCCCGCGACTTGGATTAGAGTCTCTTTTGGAATAAGCCTGAATGATCCGAGTAGCATCTCAG"
//	s := "TACAAGCCCCCTAAGCTTTCAAGGAATGAACCGACAGTAGTGCAGAGCAACTTAACGTCTGAATACGGCAAGGGCCAACCCATTTCCCTTCGTGTAATGCCGTTGCGATGAACCGCTGTGCACAATTGTTGACGTGTGCGAAACGTAGAATGTACCGAAGCCACATAAAAGAAGTTGGAAGGAGCCTAGCGAAACATTTCTAGACTTCTGGTCCTGGGCGCACGTGCCATAACAAACTGCAATGCATCTCTGCGTGTCCCTCCGTGGGCATCCCGTGCACCCACCTTATAATCGAGGAGCCTATGGGAATATGAAATACCACACCCAGCGTCGTGTCACATCAAGTAACCATGTATTACTGGATTCCCGGACCTTACCCGTGGAGTCGTCTGTTGCTCTTAATGTATAACAAGAGCTAGTAAGCTTAACCCGCGCTAGAACCACATGCTTACCGATCCGTCCCATGAGGAGGCTTAGCTAAGCCTCCTCATGGGACGGATCGGTAAGCATCAGAACTTACTACGTGACGAACCGTAAACCACTAAAGTTGTCGAGCTGCTCTTGGCTATGTCGCCAGTTTGCTACGGCTTGGAGCTCTGTTCGTGCCCTGACGTTGTATTTGACTGTGTTTTGAAGGCTAGATAATGTAGAATCTGTCAGAATGTCAAACATTCCCAGCACAGAACACGCTGCGCCCATAAGACCGTAGTATAGGTGTTGCAGTCTCAGTACAACAGATTTGGTAAGAGTCACTTCTACGATTAGCAATTGAAATTAGGTGAAAGCGACGTAGATCGGCACATTTTGTGCAGACGCGGTACGTGAAATCCACTTGTCCATCTGTAAAAAGCGGAGCGCCTCTTCTGCAAGATGGTCAGTTTAATTAAAAACTGGCGGTCCAATGTCGTTTGGGGCGCTATTTAAACCGGAAGCTGTTGGGACAATGTCTGTAGG"
	s = strings.Replace(s,"T","U",-1)
	
	uniq := make(map[string]int)
	
	r := ""
	
	for {
		if rna[s[:3]] == "M" {
			uniq[ParseORF(s)] = 1
		}
		
		switch s[0] {
			case 'A': r = "U" + r
			case 'U': r = "A" + r
			case 'G': r = "C" + r
			case 'C': r = "G" + r
		}
		
		s = s[1:]
		if len(s)<3 {break}
	}
	
	for {
		if rna[r[:3]] == "M" {
			uniq[ParseORF(r)] = 1
		}
		r = r[1:]
		if len(r)<3 {break}
	}
	
	for k := range uniq {
		fmt.Print(k)
	}
}
