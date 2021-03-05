package service

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gusgins/meli-backend/config"
)

// Service exported
type Service struct {
	Config config.Configuration
}

// NewService creates service with config
func NewService(config config.Configuration) Service {
	service := Service{config}
	return service
}

// PostMutant returns
func (s Service) PostMutant(c *gin.Context) {

	var registry Registry
	if err := c.BindJSON(&registry); err != nil {
		c.JSON(400, gin.H{"error": "invalid request: " + err.Error()})
		return
	}
	if err := registry.validate(); err != nil {
		c.JSON(400, gin.H{"error": "invalid request: " + err.Error()})
		return
	}
	if registry.isMutant() {
		c.JSON(200, gin.H{"status": "authorized"})
	} else {
		c.JSON(403, gin.H{"error": "unauthorized"})
	}
}

// Registry to check if it's mutant
type Registry struct {
	Dna  []string `json:"dna"`
	size int
	code uint64
}

// New Creates Registry from dna
func (r *Registry) validate() error {
	r.size = len(r.Dna)
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
	r.code = decode(code, "0123")
	return nil
}

func (r Registry) isMutant() bool {
	return IsMutant(r.size, r.Dna)
}

// decode Converts enc into base len(base) int
func decode(enc, base string) uint64 {
	var nb uint64
	lbase := len(base)
	le := len(enc)
	for i := 0; i < le; i++ {
		mult := 1
		for j := 0; j < le-i-1; j++ {
			mult *= lbase
		}
		nb += uint64(strings.IndexByte(base, enc[i]) * mult)
	}
	return nb
}

// RegistryError for invalid registry
type RegistryError struct {
	Err string
}

func (r RegistryError) Error() string { return r.Err }
