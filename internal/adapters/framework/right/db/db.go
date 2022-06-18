package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	// Connect to database
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("db connection failure: %v", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping failed: %v", err)
	}
	return &Adapter{db}, nil
}

func (da Adapter) CloseDbConnection() {
	err := da.db.Close()
	if err != nil {
		log.Fatalf("db ping failed: %v", err)
	}
}

func (da Adapter) AddToHistory(answer int32, operation string) error {

	queryString, args, err := sq.Insert("airth_history").Columns("date", "answer", "operation").
		Values(time.Now(), answer, operation).ToSql()

	if err != nil {
		return err
	}

	_, err = da.db.Exec(queryString, args...)
	if err != nil {
		return err
	}
	return nil

}
