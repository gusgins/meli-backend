package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gusgins/meli-backend/config"
	"github.com/gusgins/meli-backend/model"
	"github.com/gusgins/meli-backend/repository"
	"github.com/gusgins/meli-backend/repository/mysql"
	"github.com/stretchr/testify/assert"
)

func TestPostMutant(t *testing.T) {
	config := config.Configuration{
		API: config.APIConfiguration{
			Port: 8080,
		},
		Database: config.DatabaseConfiguration{
			Host:     "localhost",
			Port:     3306,
			Name:     "meli_backend",
			User:     "root",
			Password: "password",
		},
	}
	var repository repository.Repository
	repository, err := mysql.NewRepository(config)
	assert.NoError(t, err)
	service := NewService(config, repository)
	c, r := gin.CreateTestContext(httptest.NewRecorder())
	r.POST("/mutant", service.PostMutant)
	// Invalid JSON
	runTestPostMutant(t, c, r, `{"dna":"ATT","TATT","AATA","ATAA"]}`, http.StatusBadRequest, "error", "invalid request: invalid character ',' after object key")
	// Invalid array size
	runTestPostMutant(t, c, r, `{"dna":["AAAA","AAAA","AAAA","AAA"]}`, http.StatusBadRequest, "error", "invalid request: invalid matrix size")
	// Invalid character
	runTestPostMutant(t, c, r, `{"dna":["AABA","AAAA","AAAA","AAA"]}`, http.StatusBadRequest, "error", "invalid request: invalid character")
	// Not mutant (Exercise)
	runTestPostMutant(t, c, r, `{"dna":["ATGCGA","CAGTGC","TTATTT","AGACGG","GCGTCA","TCACTG"]}`, http.StatusForbidden, "error", "unauthorized")
	// Mutant (Exercise)
	runTestPostMutant(t, c, r, `{"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}`, http.StatusOK, "status", "authorized")
	// Mutant
	runTestPostMutant(t, c, r, `{"dna":["AAAA","AAAA","AAAA","AAAA"]}`, http.StatusOK, "status", "authorized")
	// Mutant by Major Diagonal (0)
	runTestPostMutant(t, c, r, `{"dna":["AAAA","TAAA","ATAT","AATA"]}`, http.StatusOK, "status", "authorized")
	// Mutant by Above Major Diagonal (4)
	runTestPostMutant(t, c, r, `{"dna":["TAGCGA","GCATGC","TTTAGT","GAGAAG","ACCCCT","TCACTG"]}`, http.StatusOK, "status", "authorized")
	// Mutant by Above Minor Diagonal (6)
	runTestPostMutant(t, c, r, `{"dna":["AGCGAT","CGTACG","TGATTT","GAAGAG","TCCCCA","GTCACT"]}`, http.StatusOK, "status", "authorized")
	// Mutant by Below Minor Diagonal (7)
	runTestPostMutant(t, c, r, `{"dna":["AGCGTT","CGTACC","TGATCT","GAACAG","TCCCCA","GCCACT"]}`, http.StatusOK, "status", "authorized")
	// Not mutant
	runTestPostMutant(t, c, r, `{"dna":["ATTT","TATT","AATA","ATAA"]}`, http.StatusForbidden, "error", "unauthorized")
}

func TestPostMutantSkipDB(t *testing.T) {
	config := config.Configuration{
		API: config.APIConfiguration{
			Port: 8080,
		},
		Database: config.DatabaseConfiguration{
			Host:     "localhost",
			Port:     3306,
			Name:     "meli_backend",
			User:     "root",
			Password: "password",
		},
	}
	var repository repository.Repository
	repository, err := mysql.NewRepository(config)
	assert.NoError(t, err)
	service := Service{config, repository, true}
	c, r := gin.CreateTestContext(httptest.NewRecorder())
	r.POST("/mutant", service.PostMutant)
	// Invalid JSON
	runTestPostMutant(t, c, r, `{"dna":"ATT","TATT","AATA","ATAA"]}`, http.StatusBadRequest, "error", "invalid request: invalid character ',' after object key")
	// Invalid array size
	runTestPostMutant(t, c, r, `{"dna":["AAAA","AAAA","AAAA","AAA"]}`, http.StatusBadRequest, "error", "invalid request: invalid matrix size")
	// Invalid character
	runTestPostMutant(t, c, r, `{"dna":["AABA","AAAA","AAAA","AAA"]}`, http.StatusBadRequest, "error", "invalid request: invalid character")
	// Not mutant (Exercise)
	runTestPostMutant(t, c, r, `{"dna":["ATGCGA","CAGTGC","TTATTT","AGACGG","GCGTCA","TCACTG"]}`, http.StatusForbidden, "error", "unauthorized")
	// Mutant (Exercise)
	runTestPostMutant(t, c, r, `{"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}`, http.StatusOK, "status", "authorized")
	// Mutant
	runTestPostMutant(t, c, r, `{"dna":["AAAA","AAAA","AAAA","AAAA"]}`, http.StatusOK, "status", "authorized")
	// Mutant by Major Diagonal (0)
	runTestPostMutant(t, c, r, `{"dna":["AAAA","TAAA","ATAT","AATA"]}`, http.StatusOK, "status", "authorized")
	// Mutant by Above Major Diagonal (4)
	runTestPostMutant(t, c, r, `{"dna":["TAGCGA","GCATGC","TTTAGT","GAGAAG","ACCCCT","TCACTG"]}`, http.StatusOK, "status", "authorized")
	// Mutant by Above Minor Diagonal (6)
	runTestPostMutant(t, c, r, `{"dna":["AGCGAT","CGTACG","TGATTT","GAAGAG","TCCCCA","GTCACT"]}`, http.StatusOK, "status", "authorized")
	// Mutant by Below Minor Diagonal (7)
	runTestPostMutant(t, c, r, `{"dna":["AGCGTT","CGTACC","TGATCT","GAACAG","TCCCCA","GCCACT"]}`, http.StatusOK, "status", "authorized")
	// Not mutant
	runTestPostMutant(t, c, r, `{"dna":["ATTT","TATT","AATA","ATAA"]}`, http.StatusForbidden, "error", "unauthorized")
}

func runTestPostMutant(t *testing.T, c *gin.Context, r *gin.Engine, jsonBody string, code int, field string, value string) {
	t.Helper()
	w := httptest.NewRecorder()
	c.Request, _ = http.NewRequest("POST", "/mutant", strings.NewReader(jsonBody))
	r.ServeHTTP(w, c.Request)
	var responseJSON map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &responseJSON)
	if err != nil {
		t.Error(err.Error())
	}
	responseData, exists := responseJSON[field]
	if !exists {
		t.Error("Invalid return")
	}
	if responseData != value {
		t.Error(fmt.Sprintf("Expected %s %s, got %s", field, value, responseData))
	}
	if w.Code != code {
		t.Error(fmt.Sprintf("Expected Code %d, got %d", code, w.Code))
	}
}

func TestGetStats(t *testing.T) {
	config := config.Configuration{
		API: config.APIConfiguration{
			Port: 8080,
		},
		Database: config.DatabaseConfiguration{
			Host:     "localhost",
			Port:     3306,
			Name:     "meli_backend",
			User:     "root",
			Password: "password",
		},
	}
	var repository repository.Repository
	repository, err := mysql.NewRepository(config)
	assert.NoError(t, err)
	service := Service{config, repository, false}
	c, r := gin.CreateTestContext(httptest.NewRecorder())
	r.GET("/stats", service.GetStats)
	t.Helper()
	w := httptest.NewRecorder()
	c.Request, _ = http.NewRequest("GET", "/stats", strings.NewReader(""))
	r.ServeHTTP(w, c.Request)

	var responseJSON model.Stats
	err = json.Unmarshal([]byte(w.Body.String()), &responseJSON)
	fmt.Println(responseJSON)
	assert.NoError(t, err)
	assert.Equal(t, w.Code, http.StatusOK)
}

func TestGetStatsError(t *testing.T) {

}
