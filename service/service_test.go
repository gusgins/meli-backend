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
			APIPort: 8080,
		},
		Database: config.DatabaseConfiguration{},
	}
	service := NewService(config)
	c, r := gin.CreateTestContext(httptest.NewRecorder())
	r.POST("/mutant", service.PostMutant)
	// JSON Inválido
	runTest(t, c, r, `{"dna":"ATT","TATT","AATA","ATAA"]}`, http.StatusBadRequest, "error", "invalid request: invalid character ',' after object key")
	// Matriz de tamaños no válidos
	runTest(t, c, r, `{"dna":["AAAA","AAAA","AAAA","AAA"]}`, http.StatusBadRequest, "error", "invalid request: Invalid matrix size")
	// No mutante (Consigna)
	runTest(t, c, r, `{"dna":["ATGCGA","CAGTGC","TTATTT","AGACGG","GCGTCA","TCACTG"]}`, http.StatusForbidden, "error", "unauthorized")
	// Mutante (Consigna)
	runTest(t, c, r, `{"dna":["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]}`, http.StatusOK, "status", "authorized")
	// Mutante
	runTest(t, c, r, `{"dna":["AAAA","AAAA","AAAA","AAAA"]}`, http.StatusOK, "status", "authorized")
	// Mutante por diagonal principal (0)
	runTest(t, c, r, `{"dna":["AAAA","TAAA","ATAT","AATA"]}`, http.StatusOK, "status", "authorized")
	// Mutante por diagonal superior a principal (4)
	runTest(t, c, r, `{"dna":["TAGCGA","GCATGC","TTTAGT","GAGAAG","ACCCCT","TCACTG"]}`, http.StatusOK, "status", "authorized")
	// Mutante por diagonal superior a inversa (6)
	runTest(t, c, r, `{"dna":["AGCGAT","CGTACG","TGATTT","GAAGAG","TCCCCA","GTCACT"]}`, http.StatusOK, "status", "authorized")
	// Mutante por diagonal inferior a inversa (7)
	runTest(t, c, r, `{"dna":["AGCGTT","CGTACC","TGATCT","GAACAG","TCCCCA","GCCACT"]}`, http.StatusOK, "status", "authorized")
	// No mutante
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
