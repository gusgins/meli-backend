package mysql

import (
	"database/sql"
	"fmt"

	// Mysql Driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/gusgins/meli-backend/config"
	"github.com/gusgins/meli-backend/model"
	repo "github.com/gusgins/meli-backend/repository"
)

type repository struct {
	db *sql.DB
}

// NewRepository Constructor for MySQL Repository
func NewRepository(c config.Configuration) (repo.Repository, error) {
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", c.Database.User, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.Name)
	db, err := sql.Open("mysql", connectionString)
	db.SetMaxIdleConns(10)
	// db.SetMaxOpenConns(0)
	if err != nil {
		return &repository{db}, err
	}

	// ping to check if db connection works
	if err = db.Ping(); err != nil {
		return &repository{db}, err
	}
	db.Exec("CREATE TABLE IF NOT EXISTS `registry`(`size` INT UNSIGNED NOT NULL,id VARBINARY(200) NOT NULL,mutant BOOLEAN NOT NULL,PRIMARY KEY (size,id)) ENGINE=MyISAM DEFAULT CHARSET=utf8;")
	return &repository{db}, err
}

func (r *repository) FindMutant(registry model.Registry) (bool, error) {
	row := r.db.QueryRow("SELECT mutant FROM registry WHERE size = ? AND id = ?", registry.Size, registry.Code)
	var mutant bool
	err := row.Scan(&mutant)
	if err == sql.ErrNoRows {
		return mutant, repo.ErrRegistryNotFound
	}
	if err != nil {
		return mutant, err
	}

	return mutant, nil
}

func (r *repository) StoreRegistry(registry model.Registry) error {
	stmt, err := r.db.Prepare("INSERT INTO registry (size, id, mutant) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(registry.Size,
		registry.Code,
		registry.Mutant,
	)
	return err
}

func (r *repository) GetStats() (model.Stats, error) {
	row := r.db.QueryRow("SELECT COUNT(NULLIF(mutant,0)) Mutants, COUNT(NULLIF(mutant,1)) Humans FROM registry")
	stats := model.Stats{}
	err := row.Scan(&stats.Mutants, &stats.Humans)
	if err == sql.ErrNoRows {
		return stats, repo.ErrStatsNotFound
	}
	return stats, nil
}

func (r *repository) Close() error {
	return r.db.Close()
}
