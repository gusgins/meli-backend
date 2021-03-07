package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRatio(t *testing.T) {
	s := Stats{
		Mutants: 10,
		Humans:  10,
	}
	assert.Equal(t, s.GetRatio(), "1.00")
}

func TestGetRatioZero(t *testing.T) {
	s := Stats{
		Mutants: 10,
		Humans:  0,
	}
	assert.Equal(t, s.GetRatio(), "0")
}
