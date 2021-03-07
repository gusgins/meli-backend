package model

import "fmt"

// Stats to store verification statistics
type Stats struct {
	Mutants int
	Humans  int
}

// GetRatio returns Stats' Mutant/Human ratio
func (s *Stats) GetRatio() string {
	if s.Humans == 0 {
		return "0"
	}
	return fmt.Sprintf("%.2f", float64(s.Mutants)/float64(s.Humans))
}
