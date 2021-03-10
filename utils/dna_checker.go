package utils

import (
	"sync"
)

type state struct {
	size           int
	dna            []string
	mutations      int
	mutant         bool
	mutationsMutex sync.Mutex
	wg             sync.WaitGroup
}
type direction struct {
	nextI int
	nextJ int
}

// IsMutant find if dna of size is mutant
func IsMutant(size int, dna []string) bool {
	s := &state{
		size:           size,
		dna:            dna,
		mutations:      0,
		mutationsMutex: sync.Mutex{},
		wg:             sync.WaitGroup{},
	}
	dRow := &direction{0, 1}
	dCol := &direction{1, 0}
	dMajD := &direction{1, 1}
	dMinD := &direction{1, -1}
	for i := 0; i < size; i++ {
		s.wg.Add(1)
		go s.check(i, 0, dRow) // Check each Row
		// dRow
		// [i=0]                  [i=1]
		// [i0][>][>][>][>][>]    [  ][ ][ ][ ][ ][ ]
		// [  ][ ][ ][ ][ ][ ]    [i0][>][>][>][>][>]
		// [  ][ ][ ][ ][ ][ ] => [  ][ ][ ][ ][ ][ ]
		// [  ][ ][ ][ ][ ][ ]    [  ][ ][ ][ ][ ][ ]
		// [  ][ ][ ][ ][ ][ ]    [  ][ ][ ][ ][ ][ ]
		// [  ][ ][ ][ ][ ][ ]    [  ][ ][ ][ ][ ][ ]

		s.wg.Add(1)
		go s.check(0, i, dCol) // Check each Column
		// dCol
		// [i=0]                  [i=1]
		// [0i][ ][ ][ ][ ][ ]    [ ][0i][ ][ ][ ][ ]
		// [v ][ ][ ][ ][ ][ ]    [ ][v ][ ][ ][ ][ ]
		// [v ][ ][ ][ ][ ][ ] => [ ][v ][ ][ ][ ][ ]
		// [v ][ ][ ][ ][ ][ ]    [ ][v ][ ][ ][ ][ ]
		// [v ][ ][ ][ ][ ][ ]    [ ][v ][ ][ ][ ][ ]
		// [v ][ ][ ][ ][ ][ ]    [ ][v ][ ][ ][ ][ ]

		if i <= size-4 { // if Diagonal has 4 or more elements
			s.wg.Add(1)
			go s.check(0, i, dMajD) // Check in Main Diagonal direction starting from [0][0] to [0][size-4]
			// dMajD - To right
			// [i=0]                  [i=1]
			// [i0][ ][ ][ ][ ][ ]    [ ][i0][ ][ ][ ][ ]
			// [  ][\][ ][ ][ ][ ]    [ ][  ][\][ ][ ][ ]
			// [  ][ ][\][ ][ ][ ] => [ ][  ][ ][\][ ][ ]
			// [  ][ ][ ][\][ ][ ]    [ ][  ][ ][ ][\][ ]
			// [  ][ ][ ][ ][\][ ]    [ ][  ][ ][ ][ ][\]
			// [  ][ ][ ][ ][ ][\]    [ ][  ][ ][ ][ ][ ]

			s.wg.Add(1)
			go s.check(i, s.size-1, dMinD) // Check in Minor Diagonal direction starting from [0][size-1] to [size-4][size-1]
			// dMinD - To bottom
			// [i=0]                  [i=1]
			// [ ][ ][ ][ ][ ][i5]    [ ][ ][ ][ ][ ][  ]
			// [ ][ ][ ][ ][/][  ]    [ ][ ][ ][ ][ ][i5]
			// [ ][ ][ ][/][ ][  ] => [ ][ ][ ][ ][/][  ]
			// [ ][ ][/][ ][ ][  ]    [ ][ ][ ][/][ ][  ]
			// [ ][/][ ][ ][ ][  ]    [ ][ ][/][ ][ ][  ]
			// [/][ ][ ][ ][ ][  ]    [ ][/][ ][ ][ ][  ]
			if i > 0 {
				s.wg.Add(1)
				go s.check(i, 0, dMajD) // Check in Major Diagonal direction starting from [1][0] to [size-4][0]
				// dMajD - To bottom
				// [i=1]                  [i=2]
				// [  ][ ][ ][ ][ ][ ]    [  ][ ][ ][ ][ ][ ]
				// [i0][ ][ ][ ][ ][ ]    [  ][ ][ ][ ][ ][ ]
				// [  ][\][ ][ ][ ][ ] => [i0][ ][ ][ ][ ][ ]
				// [  ][ ][\][ ][ ][ ]    [  ][\][ ][ ][ ][ ]
				// [  ][ ][ ][\][ ][ ]    [  ][ ][\][ ][ ][ ]
				// [  ][ ][ ][ ][\][ ]    [  ][ ][ ][\][ ][ ]

				s.wg.Add(1)
				go s.check(0, s.size-1-i, dMinD) // Check in Minor Diagonal direction starting from [size-2][0] to [3][0]
				// dMinD - To left
				// [i=1] => r=s.size-1-i  [i=2] => r=s.size-1-i
				// [ ][ ][ ][ ][r0][ ]    [ ][ ][ ][r0][ ][ ]
				// [ ][ ][ ][/][  ][ ]    [ ][ ][/][  ][ ][ ]
				// [ ][ ][/][ ][  ][ ] => [ ][/][ ][  ][ ][ ]
				// [ ][/][ ][ ][  ][ ]    [/][ ][ ][  ][ ][ ]
				// [/][ ][ ][ ][  ][ ]    [ ][ ][ ][  ][ ][ ]
				// [ ][ ][ ][ ][  ][ ]    [ ][ ][ ][  ][ ][ ]
			}
		}
	}
	s.wg.Wait()
	return s.mutant
}

func (s *state) check(startI int, startJ int, dir *direction) {
	gene := s.dna[startI][startJ]
	rep := 1
	for i, j := startI+dir.nextI, startJ+dir.nextJ; i < s.size && j < s.size && i >= 0 && j >= 0 && !s.mutant; i, j = i+dir.nextI, j+dir.nextJ {
		if gene == s.dna[i][j] {
			rep++
			if rep == 4 {
				s.mutationsMutex.Lock()
				s.mutations++
				if s.mutations > 1 {
					s.mutant = true
				}
				s.mutationsMutex.Unlock()
				gene = ' '
			}
		} else {
			gene = s.dna[i][j]
			rep = 1
		}
	}
	s.wg.Done()
}
