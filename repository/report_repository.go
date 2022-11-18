package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/sqltocsv"
)

type ReportRepository struct {
	db *sqlx.DB
}

func NewReportRepository(db *sqlx.DB) *ReportRepository {
	return &ReportRepository{db: db}
}

func (t ReportRepository) GetReport(year int, month int) (string, error) {
	createQuery := fmt.Sprintf("select * from %s where extract(year from datetime)=$1 and extract(month from datetime)=$2",
		transactionTable)
	rows, err := t.db.Query(createQuery, year, month)
	if err != nil {
		return "", err
	}
	filename := fmt.Sprintf("report-%d-%d.csv", year, month)
	err = sqltocsv.WriteFile("./reports/"+filename, rows)
	if err != nil {
		return "", err
	}
	return filename, err
}
