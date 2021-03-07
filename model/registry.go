package model

import (
	"errors"
	"fmt"
	"strings"

	"github.com/eknkc/basex"
	"github.com/gusgins/meli-backend/utils"
)

// Registry to check if it's mutant
type Registry struct {
	Dna    []string `json:"dna"`
	Size   int
	Code   string
	Mutant bool
}

// ErrInvalidCharacter is returned by Validate when invalid characters
// are found in the dna of the registry
var ErrInvalidCharacter = errors.New("invalid character")

// ErrInvalidMatrixSize is returned by Validate when matrix side is not
// consistent in the dna of the registry
var ErrInvalidMatrixSize = errors.New("invalid matrix size")

// Validate checks dna consistency
func (r *Registry) Validate() error {
	r.Size = len(r.Dna)
	values := map[rune]string{'A': "0", 'T': "1", 'C': "2", 'G': "3"}
	code := ""
	size := len(r.Dna)
	for _, s := range r.Dna {
		if len(s) != size {
			return ErrInvalidMatrixSize
		}
		for _, c := range s {
			if val, found := values[c]; found {
				code += val
			} else {
				return ErrInvalidCharacter
			}
		}
	}
	var err error
	r.Code, err = generateCode(code)
	fmt.Println(r.Dna, r.Size, r.Code)
	codeDna, err := decode(r.Size, r.Code)
	fmt.Println(codeDna)
	return err
}

// IsMutant returns if registry is mutant
func (r *Registry) IsMutant() bool {
	r.Mutant = utils.IsMutant(r.Size, r.Dna)
	return r.Mutant
}

func generateCode(code string) (string, error) {
	encoding, err := basex.NewEncoding("0123")
	if err != nil {
		return "", err
	}
	decodedCode, err := encoding.Decode(code)
	if err != nil {
		return "", ErrInvalidCharacter
	}
	return string(decodedCode), nil
}

func decode(size int, code string) ([]string, error) {
	values := map[rune]rune{'0': 'A', '1': 'T', '2': 'C', '3': 'G'}
	dna := make([]string, size, size)
	encoding, err := basex.NewEncoding("0123")
	if err != nil {
		return dna, err
	}
	encodedCode := encoding.Encode([]byte(code))
	for i := size - 1; i >= 0; i-- {
		var s string
		if len(encodedCode) >= size {
			s = encodedCode[len(encodedCode)-size:]
			encodedCode = encodedCode[0 : len(encodedCode)-size]
		} else {
			s = encodedCode[0:]
			encodedCode = encodedCode[0:0]
		}
		dna[i] = strings.Map(func(r rune) rune { return values[r] }, s)
	}
	return dna, nil
}
