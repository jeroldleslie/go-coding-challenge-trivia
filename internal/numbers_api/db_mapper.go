package numbers_api

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

type DBMapper struct {
	DB *sql.DB
}

func (dbMapper *DBMapper) fetchDetails(ctx context.Context, number string, text string) (numbers []Number, err error) {
	number = strings.TrimSpace(number)
	text = strings.TrimSpace(text)

	conditions := make([]string, 0)
	if number != "" {
		conditions = append(conditions, fmt.Sprintf("number='%s'", number))
	}

	if text != "" {
		conditions = append(conditions, fmt.Sprintf("text like '%s'", "%"+text+"%"))
	}

	finalCond := strings.Join(conditions, " and ")
	query := "select * from numbers"

	if finalCond != "" {
		query = query + " where " + finalCond
	}

	query = query + " order by number asc"

	logrus.Infof("Executed Query: %s", query)

	rows, err := dbMapper.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var number Number
		err := rows.Scan(&number.Number, &number.Type, &number.Text, &number.Found)
		if err != nil {
			fmt.Errorf("%+v", err)
		} else {
			numbers = append(numbers, number)
		}
	}

	if numbers == nil {
		err = ErrNoDataAvailable
	}
	return numbers, err
}
