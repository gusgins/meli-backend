package mysql

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gusgins/meli-backend/model"
	r "github.com/gusgins/meli-backend/repository"
	"github.com/gusgins/meli-backend/utils"
	"github.com/stretchr/testify/assert"
)

var registry = &model.Registry{
	Dna: []string{"AAAA", "AAAA", "AAAA", "AAAA"},
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestFindMutant(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer repo.Close()

	query := "SELECT mutant FROM registry WHERE size = \\? AND id = \\?"

	registry.Validate()
	registry.Mutant = utils.IsMutant(registry.Size, registry.Dna)
	rows := sqlmock.NewRows([]string{"mutant"}).
		AddRow(registry.Mutant)

	mock.ExpectQuery(query).WithArgs(registry.Size, registry.Code).WillReturnRows(rows)

	mutant, err := repo.FindMutant(*registry)
	assert.NoError(t, err)
	assert.Equal(t, mutant, registry.IsMutant())
}

func TestFindMutantNotFound(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer repo.Close()

	query := "SELECT mutant FROM registry WHERE size = \\? AND id = \\?"

	rows := sqlmock.NewRows([]string{"mutant"})

	mock.ExpectQuery(query).WithArgs(registry.Size, registry.Code).WillReturnRows(rows)

	_, err := repo.FindMutant(*registry)
	assert.EqualError(t, err, r.ErrRegistryNotFound.Error())
}

func TestFindMutantError(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer repo.Close()

	query := "SELECT mutant FROM registry WHERE size = \\? AND id = \\?"

	mock.ExpectQuery(query).WithArgs(registry.Size, registry.Code).WillReturnError(sql.ErrConnDone)

	_, err := repo.FindMutant(*registry)
	assert.Error(t, err)
}

func TestStoreRegistry(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer repo.Close()

	registry.Validate()
	registry.Mutant = utils.IsMutant(registry.Size, registry.Dna)
	query := "INSERT INTO registry \\(size, id, mutant\\) VALUES\\(\\?, \\?, \\?\\)"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(registry.Size, registry.Code, registry.Mutant).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.StoreRegistry(*registry)
	assert.NoError(t, err)
}

func TestStoreRegistryError(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer repo.Close()

	registry.Validate()
	registry.Mutant = utils.IsMutant(registry.Size, registry.Dna)
	query := "INSERT INTO registry \\(size, id, mutant\\) VALUES\\(\\?, \\?, \\?\\)"

	prep := mock.ExpectPrepare(query).WillReturnError(sql.ErrConnDone)
	prep.ExpectExec().WithArgs(registry.Size, registry.Code, registry.Mutant).WillReturnError(sql.ErrConnDone)

	err := repo.StoreRegistry(*registry)
	assert.Error(t, err)
}

func TestGetStats(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer repo.Close()

	query := "SELECT COUNT\\(NULLIF\\(mutant,0\\)\\) Mutants, COUNT\\(NULLIF\\(mutant,1\\)\\) Humans FROM registry"

	rows := sqlmock.NewRows([]string{"Mutants", "Humans"}).AddRow(0, 0)

	mock.ExpectQuery(query).WillReturnRows(rows)

	_, err := repo.GetStats()
	assert.NoError(t, err)
}

func TestGetStatsError(t *testing.T) {
	db, mock := NewMock()
	repo := &repository{db}
	defer repo.Close()

	query := "SELECT COUNT\\(NULLIF\\(mutant,0\\)\\) Mutants, COUNT\\(NULLIF\\(mutant,1\\)\\) Humans FROM registry"

	rows := sqlmock.NewRows([]string{"Mutants", "Humans"})

	mock.ExpectQuery(query).WillReturnRows(rows)

	_, err := repo.GetStats()
	assert.EqualError(t, err, r.ErrStatsNotFound.Error())
}
