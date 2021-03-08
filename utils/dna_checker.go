package utils

type state struct {
	genes             []byte
	geneStringLengths []int
	mutations         int
	mutant            bool
}

// IsMutant find if dna of size is mutant
//
// N° - Checks
//
// [0] - Major Diagonal
// 	[0][ ][ ][ ][ ][ ]
// 	[ ][0][ ][ ][ ][ ]
// 	[ ][ ][0][ ][ ][ ]
// 	[ ][ ][ ][0][ ][ ]
// 	[ ][ ][ ][ ][0][ ]
// 	[ ][ ][ ][ ][ ][0]
// [1] - Minor Diagonal
// 	[ ][ ][ ][ ][ ][1]
// 	[ ][ ][ ][ ][1][ ]
// 	[ ][ ][ ][1][ ][ ]
// 	[ ][ ][1][ ][ ][ ]
// 	[ ][1][ ][ ][ ][ ]
// 	[1][ ][ ][ ][ ][ ]
// [2] - Rows
// 	[2][2][2][2][2][2]    [ ][ ][ ][ ][ ][ ]
// 	[ ][ ][ ][ ][ ][ ]    [2][2][2][2][2][2]
// 	[ ][ ][ ][ ][ ][ ] => [ ][ ][ ][ ][ ][ ]
// 	[ ][ ][ ][ ][ ][ ]    [ ][ ][ ][ ][ ][ ]
// 	[ ][ ][ ][ ][ ][ ]    [ ][ ][ ][ ][ ][ ]
// 	[ ][ ][ ][ ][ ][ ]    [ ][ ][ ][ ][ ][ ]
// [3] - Columns
// 	[3][ ][ ][ ][ ][ ]    [ ][3][ ][ ][ ][ ]
// 	[3][ ][ ][ ][ ][ ]    [ ][3][ ][ ][ ][ ]
// 	[3][ ][ ][ ][ ][ ] => [ ][3][ ][ ][ ][ ]
// 	[3][ ][ ][ ][ ][ ]    [ ][3][ ][ ][ ][ ]
// 	[3][ ][ ][ ][ ][ ]    [ ][3][ ][ ][ ][ ]
// 	[3][ ][ ][ ][ ][ ]    [ ][3][ ][ ][ ][ ]
// [4] - Above Major Diagonal
// 	[ ][4][ ][ ][ ][ ]    [ ][ ][4][ ][ ][ ]
// 	[ ][ ][4][ ][ ][ ]    [ ][ ][ ][4][ ][ ]
// 	[ ][ ][ ][4][ ][ ] => [ ][ ][ ][ ][4][ ]
// 	[ ][ ][ ][ ][4][ ]    [ ][ ][ ][ ][ ][4]
// 	[ ][ ][ ][ ][ ][4]    [ ][ ][ ][ ][ ][ ]
// 	[ ][ ][ ][ ][ ][ ]    [ ][ ][ ][ ][ ][ ]
// [5] - Below Major Diagonal
// 	[ ][ ][ ][ ][ ][ ]    [ ][ ][ ][ ][ ][ ]
// 	[5][ ][ ][ ][ ][ ]    [ ][ ][ ][ ][ ][ ]
// 	[ ][5][ ][ ][ ][ ] => [5][ ][ ][ ][ ][ ]
// 	[ ][ ][5][ ][ ][ ]    [ ][5][ ][ ][ ][ ]
// 	[ ][ ][ ][5][ ][ ]    [ ][ ][5][ ][ ][ ]
// 	[ ][ ][ ][ ][5][ ]    [ ][ ][ ][5][ ][ ]
// [6] - Above Minor Diagonal
// 	[ ][ ][ ][ ][6][ ]    [ ][ ][ ][6][ ][ ]
// 	[ ][ ][ ][6][ ][ ]    [ ][ ][6][ ][ ][ ]
// 	[ ][ ][6][ ][ ][ ] => [ ][6][ ][ ][ ][ ]
// 	[ ][6][ ][ ][ ][ ]    [6][ ][ ][ ][ ][ ]
// 	[6][ ][ ][ ][ ][ ]    [ ][ ][ ][ ][ ][ ]
// 	[ ][ ][ ][ ][ ][ ]    [ ][ ][ ][ ][ ][ ]
// [7] - Below Minor Diagonal
// 	[ ][ ][ ][ ][ ][ ]    [ ][ ][ ][ ][ ][ ]
// 	[ ][ ][ ][ ][ ][7]    [ ][ ][ ][ ][ ][ ]
// 	[ ][ ][ ][ ][7][ ] => [ ][ ][ ][ ][ ][7]
// 	[ ][ ][ ][7][ ][ ]    [ ][ ][ ][ ][7][ ]
// 	[ ][ ][7][ ][ ][ ]    [ ][ ][ ][7][ ][ ]
// 	[ ][7][ ][ ][ ][ ]    [ ][ ][7][ ][ ][ ]
func IsMutant(size int, dna []string) bool {
	s := state{
		genes:             make([]byte, 8),
		geneStringLengths: make([]int, 8),
		mutations:         0,
	}
	s.genes[0] = ' '
	s.genes[1] = ' '
	for i := 0; i < size; i++ {
		s.checkGene(0, dna[i][i])
		if s.mutant {
			return true
		}
		s.checkGene(1, dna[i][size-1-i])
		if s.mutant {
			return true
		}
		s.genes[2] = ' '
		s.genes[3] = ' '
		for j := 0; j < size; j++ {
			s.checkGene(2, dna[i][j])
			if s.mutant {
				return true
			}
			s.checkGene(3, dna[j][i])
			if s.mutant {
				return true
			}
		}
	}
	for i := 1; i < size-3; i++ {
		s.genes[4] = ' '
		s.genes[5] = ' '
		s.genes[6] = ' '
		s.genes[7] = ' '
		for j := i; j < size; j++ {
			s.checkGene(4, dna[j-i][j])
			if s.mutant {
				return true
			}
			s.checkGene(5, dna[j][j-i])
			if s.mutant {
				return true
			}
			s.checkGene(6, dna[j-i][(size-1)-(j)])
			if s.mutant {
				return true
			}
			s.checkGene(7, dna[j][(size-1)-(j-i)])
			if s.mutant {
				return true
			}
		}
	}
	return false
}

func (s *state) checkGene(gene int, dnaGene byte) {
	if s.genes[gene] != dnaGene {
		s.genes[gene] = dnaGene
		s.geneStringLengths[gene] = 0
	}
	s.geneStringLengths[gene]++
	if s.geneStringLengths[gene] >= 4 {
		s.mutations++
		s.geneStringLengths[gene] = 0
	}
	if s.mutations > 1 {
		s.mutant = true
	}
}
