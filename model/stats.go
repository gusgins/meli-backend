package model

// Stats to store verification statistics
type Stats struct {
	Mutants int
	Humans  int
}

// Ratio returns Stats' Mutant/Human ratio
func (s *Stats) GetRatio() float64 {
	if s.Humans == 0 {
		return 0
	}
	return float64(s.Mutants) / float64(s.Humans)
}
