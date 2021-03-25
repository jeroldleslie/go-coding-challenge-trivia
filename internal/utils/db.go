package utils

import (
	"database/sql"
	"github.com/pkg/errors"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

/*func ConnectDB(dbConnectionStr string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbConnectionStr)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return db, nil
}*/

func ConnectDBWithInfiniteRetry(dbConnectionStr string, retryDuration time.Duration) (*sql.DB, error) {
	for {
		db, err := sql.Open("mysql", dbConnectionStr)
		if err == nil {
			err = db.Ping()
			if err == nil {
				return db, nil
			}
		}
		log.Printf("retrying database connection in %f seconds...", retryDuration.Seconds())
		time.Sleep(retryDuration)
	}
	return nil, errors.New("Database not connected")
}
