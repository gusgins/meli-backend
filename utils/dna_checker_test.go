package utils

import (
	"math/rand"
	"testing"
)

func TestIsMutant(t *testing.T) {
	got := isMutant([]string{
		"ATGCGA",
		"CAGTGC",
		"TTATGT",
		"AGAAGG",
		"CCCCTA",
		"TCACTG",
	})
	if !got {
		t.Error("isMutant('ATGCGA','CAGTGC','TTATGT','AGAAGG','CCCCTA','TCACTG') = false; want true")
	}
	got = isMutant([]string{
		"AAAACACGCA",
		"ACTGCGAGGA",
		"TGCTGTATTA",
		"TATGCAAGTG",
		"GGGGCAGTCT",
		"CCTAGACGAG",
		"ACAGAACACG",
		"GGATGGGGGT",
		"TTCAACGCTA",
		"TTAGACGATA"})
	if !got {
		t.Error("isMutant() = false; want true")
	}
	got = isMutant([]string{
		"AATTGTTGCG",
		"TGGCTTCGAT",
		"CAGCTATTTT",
		"TTCGGCGGAG",
		"GCTGTCTCAT",
		"CAGTGCCATA",
		"CCAAGCATAT",
		"CCATCAACCA",
		"GCGGACTTAG",
		"AGGGAGCGAT"})
	if !got {
		t.Error("isMutant() = false; want true")
	}
	got = isMutant([]string{
		"TGGC",
		"CAGG",
		"GACA",
		"CGGA"})
	if got {
		t.Error("isMutant() = true; want false")
	}
	got = isMutant([]string{
		"GTTCC",
		"CGACA",
		"ATTTC",
		"CACCA",
		"TTATA"})
	if got {
		t.Error("isMutant() = true; want false")
	}
	got = isMutant([]string{
		"TTGAGG",
		"TACACA",
		"ACCACT",
		"AGCGTG",
		"AACTCT",
		"CCCTTC"})
	if got {
		t.Error("isMutant() = true; want false")
	}
	got = isMutant([]string{
		"GGCGCTTGT",
		"CTTCAATAC",
		"CTGAGACTG",
		"CAAAGACGT",
		"CGCCTTATT",
		"ATGTGTTTC",
		"GCCACATAT",
		"TCCATGGTA",
		"ATGGACCTG"})
	if !got {
		t.Error("isMutant() = false; want true")
	}
	got = isMutant([]string{
		"CGGGGCTT",
		"TCAATGGC",
		"ATGCATGA",
		"GATCTCTC",
		"TATCGCGA",
		"GAGACCTC",
		"AATGCTAT",
		"GTTGTACA"})
	if !got {
		t.Error("isMutant() = false; want true")
	}
	got = isMutant([]string{
		"CATTGACC",
		"AGAGTGAC",
		"TTAAAGAA",
		"GATTCGGC",
		"TGCTTCGA",
		"CTAAGGCG",
		"TAGCAAAT",
		"ACTAGAAT"})
	if got {
		t.Error("isMutant() = true; want false")
	}
	got = isMutant([]string{
		"ACATACCA",
		"GTGTTACA",
		"GGCATAAG",
		"ATGGGCTC",
		"GTGCCGTA",
		"AAGGGGAG",
		"ATGATGGG",
		"TCCTTCCT"})
	if !got {
		t.Error("isMutant() = false; want true")
	}
	got = isMutant([]string{
		"TATAGAACC",
		"AGCATAGGC",
		"TCGTTACTG",
		"GTGAAATAG",
		"AGCATGGCG",
		"AGTCGGCAT",
		"TGGTTGGCC",
		"TGGTTCATT",
		"GAGCACAGA"})
	if !got {
		t.Error("isMutant() = false; want true")
	}
	got = isMutant([]string{
		"CGCACT",
		"AGAGCG",
		"GGTACG",
		"TAGTAT",
		"TTCACG",
		"CCCGTA"})
	if got {
		t.Error("isMutant() = true; want false")
	}
	got = isMutant([]string{
		"GGCCC",
		"ATGCG",
		"TTTAA",
		"GAGAC",
		"CCAGA"})
	if got {
		t.Error("isMutant() = true; want false")
	}
	got = isMutant([]string{
		"TGGAGG",
		"TAAAGG",
		"CCGAGC",
		"TTGGGC",
		"GTCTTC",
		"GAACCA"})
	if got {
		t.Error("isMutant() = true; want false")
	}
	got = isMutant([]string{
		"ATGAA",
		"AGCTC",
		"GAATC",
		"AGGGT",
		"TGCCT"})
	if got {
		t.Error("isMutant() = true; want false")
	}
	got = isMutant([]string{
		"ATGTTCA",
		"AACAAAT",
		"ACAAACT",
		"GCCATCA",
		"TGAGCTG",
		"CTTATCA",
		"AGGTCAG"})
	if !got {
		t.Error("isMutant() = false; want true")
	}
	got = isMutant([]string{
		"GAAGCGT",
		"GAAGAAG",
		"TTTATGC",
		"TAACAAC",
		"CTTGGCC",
		"TGTTTAA",
		"GAACATG"})
	if got {
		t.Error("isMutant() = true; want false")
	}
	got = isMutant([]string{
		"AGTCA",
		"ACTTG",
		"ATCAA",
		"TTCCC",
		"CATAC"})
	if got {
		t.Error("isMutant() = true; want false")
	}
	got = isMutant([]string{
		"AACGAGGGGC",
		"TTTTTTCAGA",
		"TGGATGTGAG",
		"CATCCATGAC",
		"CGGACAGGCA",
		"CATGTCTATT",
		"ATGCTGGATC",
		"AATTCAGTAT",
		"GCGCGAAGGA",
		"CAACGTTTAC"})
	if !got {
		t.Error("isMutant() = false; want true")
	}
	got = isMutant([]string{
		"GACTCTT",
		"TAGTGTG",
		"CGGCAAA",
		"ACTCAAA",
		"TTTTCTA",
		"CCAGGCT",
		"ACGCGTC"})
	if !got {
		t.Error("isMutant() = false; want true")
	}
	got = isMutant([]string{
		"TTTA",
		"CTAG",
		"CCAC",
		"AGTC"})
	if got {
		t.Error("isMutant() = true; want false")
	}
	got = isMutant([]string{
		"AAAGACGC",
		"CGAAGCGA",
		"GATCAGCA",
		"CCCATACA",
		"CAGAGATC",
		"TGACTACT",
		"GGACCGGT",
		"TTATAACT"})
	if got {
		t.Error("isMutant() = true; want false")
	}
	got = isMutant([]string{
		"AAAACTTGT",
		"TCGGATGTC",
		"AACTTATAC",
		"CATAAAAGC",
		"GGATTTAGG",
		"GGCTCTGTG",
		"CTTCAGTTA",
		"TGAACCCGA",
		"CAGGCGCCA"})
	if !got {
		t.Error("isMutant() = false; want true")
	}
	got = isMutant([]string{
		"CGTACA",
		"AATGTA",
		"CTTTGA",
		"AGCATC",
		"TCGGCT",
		"ACTCGG"})
	if got {
		t.Error("isMutant() = true; want false")
	}
	got = isMutant([]string{
		"AACAA",
		"CAACA",
		"ACAAC",
		"AACAA",
		"AAAGA"})
	if !got {
		t.Error("isMutant() = false; want true")
	}
	got = isMutant([]string{
		"AACAA",
		"CAACA",
		"ACAAC",
		"AACAA",
		"AAAGA"})
	if !got {
		t.Error("isMutant() = false; want true")
	}
	got = isMutant([]string{
		"AACAA",
		"CAACA",
		"AAAGC",
		"ACCAA",
		"AGAGA"})
	if !got {
		t.Error("isMutant() = false; want true")
	}
	got = isMutant([]string{
		"AACCA",
		"CACCA",
		"AGAAC",
		"AGAAA",
		"AAAGA"})
	if !got {
		t.Error("isMutant() = false; want true")
	}
	got = isMutant([]string{
		"AACAA",
		"CAACA",
		"ACAAC",
		"AACAA",
		"AAAGA"})
	if !got {
		t.Error("isMutant() = false; want true")
	}
	got = isMutant([]string{
		"TGTT",
		"TTAC",
		"AAAG",
		"GGCC"})
	if got {
		t.Error("isMutant() = true; want false")
	}
	got = isMutant([]string{
		"GTCCAGAGTG",
		"CCATCATGCT",
		"TGCAACCTCA",
		"ACGCCCCGAA",
		"TGGCCCAATT",
		"CGCTAGCCAT",
		"TGCCCGTTTA",
		"CAGATTCGTC",
		"GGGGGGGTGA",
		"CCGTAGCGGT"})
	if !got {
		t.Error("isMutant() = false; want true")
	}
	got = isMutant([]string{
		"ATGCTT",
		"ATCGGT",
		"AAACGG",
		"CATTCG",
		"AGATTT",
		"GTATGC"})
	if got {
		t.Error("isMutant() = true; want false")
	}
	got = isMutant([]string{
		"GTCCTTCC",
		"ATTTACTC",
		"GGGGATAG",
		"AGCAGGTG",
		"AGTAGGGG",
		"GTTGCACC",
		"ACAGCACA",
		"CCTCATCA"})
	if !got {
		t.Error("isMutant() = false; want true")
	}
	got = isMutant([]string{
		"TCCCT",
		"TGGCA",
		"TTTGA",
		"TGATA",
		"GCATG"})
	if got {
		t.Error("isMutant() = true; want false")
	}
	got = isMutant([]string{
		"ACAAT",
		"CGGGG",
		"GGTCA",
		"TGCAG",
		"TGGAT"})
	if !got {
		t.Error("isMutant() = false; want true")
	}
	got = isMutant([]string{
		"TAAGTT",
		"CTCCAA",
		"CCGGTG",
		"GGCCCT",
		"GGCCGC",
		"GGCGCA"})
	if got {
		t.Error("isMutant() = true; want false")
	}
	got = isMutant([]string{
		"GGTCTTGGTA",
		"CGCAATGCAC",
		"GTCAGATGAA",
		"ACGTTACGCG",
		"AGTGCTCACG",
		"CATACCTGCC",
		"TCCGTAAACT",
		"TACGTCTGAC",
		"CAGAACTGTC",
		"TCTTCCAATT"})
	if got {
		t.Error("isMutant() = true; want false")
	}
	got = isMutant([]string{
		"CGCT",
		"TGGG",
		"CTCC",
		"CTTG"})
	if got {
		t.Error("isMutant() = true; want false")
	}
	got = isMutant([]string{
		"GTAAA",
		"ATCCG",
		"ATATA",
		"GGCCA",
		"CGGAT"})
	if got {
		t.Error("isMutant() = true; want false")
	}
	got = isMutant([]string{
		"TCGAGCGTT",
		"GTCATAGGG",
		"CTTATATTT",
		"GAAACTCAT",
		"CCTACACAG",
		"ACGCTAGGG",
		"AGTCTCATC",
		"TATTGAGAC",
		"GGTTAAATG"})
	if got {
		t.Error("isMutant() = true; want false")
	}
	got = isMutant([]string{
		"TACTTAGC",
		"CCATATGT",
		"GACGGGAC",
		"ACTGGTCT",
		"TTCCTAGA",
		"AGGATCCT",
		"TCGACGCA",
		"TACGCAAT"})
	if got {
		t.Error("isMutant() = true; want false")
	}
	got = isMutant([]string{
		"TAAA",
		"TAAC",
		"TACC",
		"CGCC"})
	if got {
		t.Error("isMutant() = true; want false")
	}
	got = isMutant([]string{
		"TTGTCG",
		"AATAGC",
		"ACGGAA",
		"GCGGAA",
		"CGGGGA",
		"ATCAGA"})
	if !got {
		t.Error("isMutant() = false; want true")
	}
	got = isMutant([]string{
		"GCTTACGT",
		"GTAGGAAC",
		"TGCAATTG",
		"AGGGGGTC",
		"AGCGAGCG",
		"TCATTGTT",
		"TGAATTCC",
		"ACGTCTGA"})
	if !got {
		t.Error("isMutant() = false; want true")
	}
	got = isMutant([]string{
		"AACGTTTG",
		"CTTAGAAT",
		"GGGATGCC",
		"TGGCGATA",
		"TGTCACGC",
		"TGGCAGCA",
		"TAGACAAG",
		"GCTGCACC"})
	if !got {
		t.Error("isMutant() = false; want true")
	}
	got = isMutant([]string{
		"TTAA",
		"TCAT",
		"GGTG",
		"TCCG"})
	if got {
		t.Error("isMutant() = true; want false")
	}
	got = isMutant([]string{
		"AGTCTTA",
		"AGGGTAT",
		"GTGCGTG",
		"AGCTGAT",
		"TCGCTTT",
		"GCGCACG",
		"CCCAAGA"})
	if got {
		t.Error("isMutant() = true; want false")
	}
	got = isMutant([]string{
		"GCGTCTGCAC",
		"GCGATAGAGG",
		"CGCAATGTTC",
		"CAAGTAAATT",
		"ACTTAAAAGT",
		"TCCCCGGTCC",
		"CGCCTGCACC",
		"GCATTAGTTT",
		"TTCCTACTAG",
		"AATGGACACT"})
	if !got {
		t.Error("isMutant() = false; want true")
	}
	got = isMutant([]string{
		"CCTGATGGTG",
		"CTCCTTGTAG",
		"ATGAACCTGC",
		"GATTGGAGTT",
		"TGGTACAGCG",
		"GAAGTATTTC",
		"CGGGATGTTC",
		"AGTTCAGAAA",
		"TAGGAAGTGG",
		"TGCTAAACTT"})
	if got {
		t.Error("isMutant() = true; want false")
	}
	got = isMutant([]string{
		"ATTGCCT",
		"CATAACC",
		"GATATGC",
		"TGATGGC",
		"AAAAGTT",
		"GTATAGA",
		"TCTGAGC"})
	if !got {
		t.Error("isMutant() = false; want true")
	}
	got = isMutant([]string{
		"CGGAAAA",
		"GCCGGGG",
		"AACAACC",
		"AGAAGAG",
		"TGCAGGA",
		"CCGATAT",
		"TAAGGTA"})
	if !got {
		t.Error("isMutant() = false; want true")
	}
	got = isMutant([]string{
		"GCTAT",
		"TTGAA",
		"ATACA",
		"TAGCC",
		"CCATA"})
	if got {
		t.Error("isMutant() = true; want false")
	}
	got = isMutant([]string{
		"GCCC",
		"CCGG",
		"TGTA",
		"GTCC"})
	if got {
		t.Error("isMutant() = true; want false")
	}
	got = isMutant([]string{
		"CTCATA",
		"ACCGCG",
		"CTCGAG",
		"CTCCTG",
		"TGACCA",
		"CGTGCA"})
	if !got {
		t.Error("isMutant() = false; want true")
	}
	got = isMutant([]string{
		"CCACTT",
		"TTGACG",
		"ATGCAC",
		"TAGGTG",
		"TAAAAT",
		"CGTTGC"})
	if got {
		t.Error("isMutant() = true; want false")
	}
	got = isMutant([]string{
		"GTGCGGCC",
		"CAATCCCG",
		"GGGATTGT",
		"CCGCTTAA",
		"CTCGGAAC",
		"AGTTTGCG",
		"AGCCAATG",
		"TTACGACT"})
	if got {
		t.Error("isMutant() = true; want false")
	}
}

func BenchmarkIsMutant(b *testing.B) {
	genes := []string{"A", "T", "C", "G"}
	// rand.Seed(time.Now().UnixNano()) // Random seed to benchmark with different values every time
	rand.Seed(1) // Constant seed to benchmark with same values every time
	dnas := make([][]string, b.N)
	for i := 0; i < b.N; i++ {
		size := rand.Intn(10) + 4
		dna := make([]string, size)
		for j := 0; j < size; j++ {
			str := ""
			for k := 0; k < size; k++ {
				str += genes[rand.Intn(4)]
			}
			dna[j] = str
		}
		dnas[i] = dna
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsMutant(len(dnas[i]), dnas[i])
		/*
			fmt.Print("DNA: ")
			fmt.Print(dnas[i])
			fmt.Print(" - Longitud: ")
			fmt.Println(len(dnas[i]))
			esMutante := IsMutant(len(dnas[i]), dnas[i])
			if esMutante {
				fmt.Println("Es mutante")
			} else {
				fmt.Println("No es mutante")
			}
		*/
	}
}
func isMutant(dna []string) bool {
	return IsMutant(len(dna), dna)
}
