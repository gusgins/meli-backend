package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gusgins/meli-backend/config"
	"github.com/gusgins/meli-backend/model"
	"github.com/gusgins/meli-backend/storage"
)

// Service exported
type Service struct {
	Config  config.Configuration
	Storage storage.Storage
}

// NewService creates service with config
func NewService(config config.Configuration, storage storage.Storage) Service {
	service := Service{config, storage}
	return service
}

// PostMutant handles mutant search
func (s Service) PostMutant(c *gin.Context) {

	var registry model.Registry
	if err := c.BindJSON(&registry); err != nil {
		c.JSON(400, gin.H{"error": "invalid request: " + err.Error()})
		return
	}
	if err := registry.Validate(); err != nil {
		c.JSON(400, gin.H{"error": "invalid request: " + err.Error()})
		return
	}

	// If err is nil, then registry was found in storage
	if isMutant, err := s.Storage.Find(registry); err == nil {
		if isMutant {
			c.JSON(200, gin.H{"status": "authorized"})
		} else {
			c.JSON(403, gin.H{"error": "unauthorized"})
		}
		return
	}

	registry.IsMutant()
	s.Storage.Store(registry)
	if registry.Mutant {
		c.JSON(200, gin.H{"status": "authorized"})
	} else {
		c.JSON(403, gin.H{"error": "unauthorized"})
	}
}

// GetStats returns db stats
func (s Service) GetStats(c *gin.Context) {

	stats, err := s.Storage.GetStats()
	if err != nil {
		c.JSON(403, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"count_mutant_dna": stats.Mutants,
		"count_human_dna":  stats.Humans,
		"ratio":            stats.GetRatio(),
	})
}
