package model

import (
	"fmt"

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

// Validate checks dna consistency
func (r *Registry) Validate() error {
	r.Size = len(r.Dna)
	values := map[rune]string{'A': "0", 'T': "1", 'C': "2", 'G': "3"}
	code := ""
	size := len(r.Dna)
	for i, s := range r.Dna {
		if len(s) != size {
			return RegistryError{Err: "invalid matrix size"}
		}
		for j, c := range s {
			if val, found := values[c]; found {
				code += val
			} else {
				return RegistryError{Err: fmt.Sprintf("invalid character %s at [%d][%d]", string(c), i, j)}
			}
		}
	}
	encoding, err := basex.NewEncoding("0123")
	if err != nil {
		return RegistryError{Err: fmt.Sprintf("error on creating Encoding 0123")}
	}
	decodedCode, err := encoding.Decode(code)
	if err != nil {
		return RegistryError{Err: fmt.Sprintf("error on decoding string %s", code)}
	}
	r.Code = string(decodedCode)
	return nil
}

// IsMutant returns if registry is mutant
func (r *Registry) IsMutant() bool {
	r.Mutant = utils.IsMutant(r.Size, r.Dna)
	return r.Mutant
}

// RegistryError for invalid registry
type RegistryError struct {
	Err string
}

func (r RegistryError) Error() string { return r.Err }
