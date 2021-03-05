package service

type estado struct {
	genes          []byte
	genesRepetidos []int
	mutaciones     int
	mutante        bool
}

func isMutant(dna []string) bool {
	return IsMutant(len(dna), dna)
}

// IsMutant find if dna of size {size} is mutant
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
	e := estado{
		genes:          make([]byte, 8),
		genesRepetidos: make([]int, 8),
		mutaciones:     0,
	}
	e.genes[0] = ' '
	e.genes[1] = ' '
	for i := 0; i < size; i++ {
		e.checkGene(0, dna[i][i])
		if e.mutante {
			return true
		}
		e.checkGene(1, dna[i][size-1-i])
		if e.mutante {
			return true
		}
		e.genes[2] = ' '
		e.genes[3] = ' '
		for j := 0; j < size; j++ {
			e.checkGene(2, dna[i][j])
			if e.mutante {
				return true
			}
			e.checkGene(3, dna[j][i])
			if e.mutante {
				return true
			}
		}
	}
	for i := 1; i < size-4; i++ {
		e.genes[4] = ' '
		e.genes[5] = ' '
		e.genes[6] = ' '
		e.genes[7] = ' '
		for j := i; j < size; j++ {
			e.checkGene(4, dna[j-1][j])
			if e.mutante {
				return true
			}
			e.checkGene(5, dna[j][j-i])
			if e.mutante {
				return true
			}
			e.checkGene(6, dna[j-1][(size-1)-(j)])
			if e.mutante {
				return true
			}
			e.checkGene(7, dna[j][(size-1)-(j-i)])
			if e.mutante {
				return true
			}
		}
	}
	return false
}

func (e *estado) checkGene(gene int, dnaGene byte) {
	if e.genes[gene] != dnaGene {
		e.genes[gene] = dnaGene
		e.genesRepetidos[gene] = 0
	}
	e.genesRepetidos[gene]++
	if e.genesRepetidos[gene] >= 4 {
		e.mutaciones++
		e.genesRepetidos[gene] = 0
	}
	if e.mutaciones > 1 {
		e.mutante = true
	}
}
