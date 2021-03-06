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
	"github.com/gusgins/meli-backend/storage"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

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
	storage := storage.NewMySQLStorage(config)
	service := NewService(config, storage)
	c, r := gin.CreateTestContext(httptest.NewRecorder())
	r.POST("/mutant", service.PostMutant)
	// Invalid JSON
	runTest(t, c, r, `{"dna":"ATT","TATT","AATA","ATAA"]}`, http.StatusBadRequest, "error", "invalid request: invalid character ',' after object key")
	// Invalid array size
	runTest(t, c, r, `{"dna":["AAAA","AAAA","AAAA","AAA"]}`, http.StatusBadRequest, "error", "invalid request: invalid matrix size")
	// Invalid character
	runTest(t, c, r, `{"dna":["AABA","AAAA","AAAA","AAA"]}`, http.StatusBadRequest, "error", "invalid request: invalid character B at [0][2]")
	// Not mutant (Exercise)
	runTest(t, c, r, `{"dna":["ATGCGA","CAGTGC","TTATTT","AGACGG","GCGTCA","TCACTG"]}`, http.StatusForbidden, "error", "unauthorized")
	// Mutant (Exercise)
	runTest(t, c, r, `{"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}`, http.StatusOK, "status", "authorized")
	// Mutant
	runTest(t, c, r, `{"dna":["AAAA","AAAA","AAAA","AAAA"]}`, http.StatusOK, "status", "authorized")
	// Mutant by Major Diagonal (0)
	runTest(t, c, r, `{"dna":["AAAA","TAAA","ATAT","AATA"]}`, http.StatusOK, "status", "authorized")
	// Mutant by Above Major Diagonal (4)
	runTest(t, c, r, `{"dna":["TAGCGA","GCATGC","TTTAGT","GAGAAG","ACCCCT","TCACTG"]}`, http.StatusOK, "status", "authorized")
	// Mutant by Above Minor Diagonal (6)
	runTest(t, c, r, `{"dna":["AGCGAT","CGTACG","TGATTT","GAAGAG","TCCCCA","GTCACT"]}`, http.StatusOK, "status", "authorized")
	// Mutant by Below Minor Diagonal (7)
	runTest(t, c, r, `{"dna":["AGCGTT","CGTACC","TGATCT","GAACAG","TCCCCA","GCCACT"]}`, http.StatusOK, "status", "authorized")
	// Not mutant
	runTest(t, c, r, `{"dna":["ATTT","TATT","AATA","ATAA"]}`, http.StatusForbidden, "error", "unauthorized")
}

func runTest(t *testing.T, c *gin.Context, r *gin.Engine, jsonBody string, code int, field string, value string) {
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
