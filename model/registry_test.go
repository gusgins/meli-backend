package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateCode(t *testing.T) {
	r := Registry{Dna: []string{
		"AACAA",
		"CAACA",
		"ACAAC",
		"AACAA",
		"AAAGA",
	}}
	err := r.Validate()
	assert.NoError(t, err)
	decodeDna, err := decode(r.Size, r.Code)
	assert.NoError(t, err)
	assert.Equal(t, r.Dna, decodeDna)
}
func TestGenerateCodeInvalid(t *testing.T) {
	r := Registry{Dna: []string{
		"AAAAA",
		"AAAKA",
		"AAAAA",
		"AAAAA",
		"AAAAA",
	}}
	err := r.Validate()
	assert.EqualError(t, err, ErrInvalidCharacter.Error())
	decodeDna, err := decode(r.Size, r.Code)
	assert.NoError(t, err)
	assert.Equal(t, []string{"", "", "", "", ""}, decodeDna)
}

func TestGenerateCodeZero(t *testing.T) {
	r := Registry{Dna: []string{
		"AAAAA",
		"AAAAA",
		"AAAAA",
		"AAAAA",
		"AAAAA",
	}}
	err := r.Validate()
	assert.NoError(t, err)
	decodeDna, err := decode(r.Size, r.Code)
	assert.NoError(t, err)
	assert.Equal(t, r.Dna, decodeDna)
}

func TestValidate(t *testing.T) {
	r := Registry{Dna: []string{
		"AACAA",
		"CAACA",
		"ACAAC",
		"AACAA",
		"AAAGA",
	}}
	err := r.Validate()
	assert.NoError(t, err)
	assert.Equal(t, r.Size, 5)
	assert.Equal(t, r.Code, "\x00\x00 \x82\b \x80\f")
}

func TestValidateErrorInvalidMatrix(t *testing.T) {
	r := Registry{Dna: []string{
		"AACAA",
		"CAAC",
		"ACAAC",
		"AACAA",
		"AAAGA",
	}}
	err := r.Validate()
	assert.EqualError(t, err, ErrInvalidMatrixSize.Error())
}

func TestValidateErrorInvalidCharacter(t *testing.T) {
	r := Registry{Dna: []string{
		"AACAA",
		"CAACZ",
		"ACAAC",
		"AACAA",
		"AAAGA",
	}}
	err := r.Validate()
	assert.EqualError(t, err, ErrInvalidCharacter.Error())
}
