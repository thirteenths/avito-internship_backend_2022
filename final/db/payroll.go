package db

import (
	"database/sql"

	"github.com/thirteenths/final/models"
)

func (db Database) GetAllPayrolles() (*models.PayrollList, error) {
	list := &models.PayrollList{}

	rows, err := db.Conn.Query("SELECT * FROM payroll")
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var payroll models.Payroll
		err := rows.Scan(&payroll.ID, &payroll.IdBalance, &payroll.Summa, &payroll.DatePayroll)
		if err != nil {
			return list, err
		}
		list.Payrolles = append(list.Payrolles, payroll)
	}
	return list, nil
}

func (db Database) AddPayroll(payroll *models.Payroll) error {
	var id int
	var datepayroll string
	query := `INSERT INTO payroll (id_balance, summa) VALUES ($1, $2) RETURNING id, date_payroll`
	err := db.Conn.QueryRow(query, payroll.IdBalance, payroll.Summa).Scan(&id, &datepayroll)
	if err != nil {
		return err
	}

	payroll.ID = id
	payroll.DatePayroll = datepayroll
	return nil
}

func (db Database) GetPayrollById(payrollId int) (models.Payroll, error) {
	payroll := models.Payroll{}

	query := `SELECT * FROM payroll WHERE id = $1;`
	row := db.Conn.QueryRow(query, payrollId)
	switch err := row.Scan(&payroll.ID, &payroll.IdBalance, &payroll.Summa, &payroll.DatePayroll); err {
	case sql.ErrNoRows:
		return payroll, ErrNoMatch
	default:
		return payroll, err
	}
}

func (db Database) DeletePayroll(payrollId int) error {
	query := `DELETE FROM payroll WHERE id = $1;`
	_, err := db.Conn.Exec(query, payrollId)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}
/*
func (db Database) UpdatePayroll(payrollId int, payrollData models.Payroll) (models.Payroll, error) {
	payroll := models.Payroll{}
	query := `UPDATE payroll SET summa=$1 WHERE id=$2 RETURNING id, id_user, summa, data_upgrate`
	err := db.Conn.QueryRow(query, balanceData.IdUser, balanceData.Summa, balanceId).Scan(&balance.ID, &balance.IdUser, &balance.Summa, &balance.DateUpgrate)
	if err != nil {
		if err == sql.ErrNoRows {
			return payroll, ErrNoMatch
		}
		return payroll, err
	}

	return payroll, nil
}*/
