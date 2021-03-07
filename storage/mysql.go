package storage

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	// Mysql Driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/gusgins/meli-backend/config"
	"github.com/gusgins/meli-backend/model"
)

type mySQLStorage struct {
	db               *sql.DB
	connectionString string
}

// NewMySQLStorage Constructor mySQLStorage
func NewMySQLStorage(c config.Configuration) Storage {
	mySQLStorage := &mySQLStorage{
		connectionString: fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", c.Database.User, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.Name),
	}
	mySQLStorage.initDatabase()
	return Storage(mySQLStorage)
}

func (s *mySQLStorage) Find(r model.Registry) (bool, error) {
	query := s.executeSQL("SELECT mutant FROM registry WHERE size=? AND id=?", []interface{}{r.Size, r.Code})
	if len(query) == 0 {
		return false, StorageError{Err: "Not Found"}
	}
	c := query[0].(map[string]interface{})
	return c["mutant"].(int64) > 0, nil
}

func (s *mySQLStorage) Store(r model.Registry) error {
	query := s.executeSQL("INSERT INTO registry (size, id, mutant) VALUES(?,?,?);", []interface{}{r.Size, r.Code, r.Mutant})
	if len(query) == 0 {
		return StorageError{Err: "Not Found"}
	}
	return nil
}

func (s *mySQLStorage) GetStats() (model.Stats, error) {
	query := s.executeSQL("SELECT COUNT(NULLIF(mutant,0)) mutants, COUNT(NULLIF(mutant,1)) humans FROM registry;", []interface{}{})
	stats := model.Stats{}
	if len(query) == 0 {
		return stats, StorageError{Err: "Not Found"}
	}
	row := query[0].(map[string]interface{})
	var err error = nil
	stats.Mutants, err = strconv.Atoi(row["mutants"].(string))
	if err != nil {
		return stats, err
	}
	stats.Humans, err = strconv.Atoi(row["humans"].(string))
	if err != nil {
		return stats, err
	}

	return stats, nil
}

func (s *mySQLStorage) initDatabase() {
	var err error
	s.db, err = sql.Open("mysql", s.connectionString)
	if err != nil {
		log.Fatal("Error while opening database connection:", err.Error())
	}
	// try to connect to db, retry connection until it works
	tries := 0
	for err = s.db.Ping(); err != nil || tries > 10; err = s.db.Ping() {
		s.db, err = sql.Open("mysql", s.connectionString)
		log.Fatal("Error on Ping to database connection:", err.Error())
		tries++
	}
	if err != nil {
		log.Fatal("Max retries exceeded when connecting to database:", err.Error())
	}
	s.executeSQL("CREATE TABLE IF NOT EXISTS `registry`(`size` INT UNSIGNED NOT NULL,id VARBINARY(200) NOT NULL,mutant BOOLEAN NOT NULL,PRIMARY KEY (size,id)) ENGINE=MyISAM DEFAULT CHARSET=utf8;", []interface{}{})
}

func (s *mySQLStorage) Clean() error {
	return s.db.Close()
}

func (s *mySQLStorage) executeSQL(queryStr string, args []interface{}) []interface{} {
	conn, err := sql.Open("mysql", s.connectionString)
	if err != nil {
		//log.Fatal("Error while opening database connection:", err.Error())
	}

	for err = conn.Ping(); err != nil; err = conn.Ping() {
		conn, err = sql.Open("mysql", s.connectionString)
		//log.Fatal("Error on Ping to database connection:", err.Error())
	}
	defer conn.Close()

	rows, err := conn.Query(queryStr, args...)
	if err != nil {
		//log.Fatal("Query failed:", err.Error())
	}
	defer rows.Close()

	columns, _ := rows.Columns()
	count := len(columns)

	var v struct {
		Data []interface{} // `json:"data"`
	}

	for rows.Next() {
		values := make([]interface{}, count)
		valuePtrs := make([]interface{}, count)
		for i := range columns {
			valuePtrs[i] = &values[i]
		}
		if err := rows.Scan(valuePtrs...); err != nil {
			//log.Fatal(err)
		}

		m := make(map[string]interface{})
		for i := range columns {
			if str, ok := values[i].(string); ok {
				m[columns[i]] = str
			} else if str, ok := values[i].(bool); ok {
				m[columns[i]] = str
			} else if str, ok := values[i].(int64); ok {
				m[columns[i]] = str
			} else if str, ok := values[i].(int); ok {
				m[columns[i]] = str
			} else if str, ok := values[i].([]uint8); ok {
				m[columns[i]] = string(str)
			} else if str, ok := values[i].(time.Time); ok {
				m[columns[i]] = str
			} else if values[i] == nil {
				m[columns[i]] = nil
			} else {
				m[columns[i]] = values[i].(string)
			}
		}
		v.Data = append(v.Data, m)
	}
	return v.Data
}

type StorageError struct {
	Err string
}

func (r StorageError) Error() string { return r.Err }
