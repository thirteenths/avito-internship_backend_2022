package db

import (
	"database/sql"

	"github.com/thirteenths/final/models"
)

func (db Database) GetAllReserves() (*models.ReserveList, error) {
	list := &models.ReserveList{}

	rows, err := db.Conn.Query("SELECT * FROM reserve")
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var reserve models.Reserve
		err := rows.Scan(&reserve.ID, &reserve.IdBalance, &reserve.Summa, &reserve.DateReserve, &reserve.IdOrder, &reserve.IdService)
		if err != nil {
			return list, err
		}
		list.Reserves = append(list.Reserves, reserve)
	}
	return list, nil
}

func (db Database) AddReserve(reserve *models.Reserve) error {
	var id int
	var dateReserve string
	query := `INSERT INTO reserve (id_balance, summa, id_order, id_service) VALUES ($1, $2, $3, $4) RETURNING id, date_reserve`
	err := db.Conn.QueryRow(query, reserve.IdBalance, reserve.Summa).Scan(&id, &dateReserve)
	if err != nil {
		return err
	}

	reserve.ID = id
	reserve.DateReserve = dateReserve
	return nil
}

func (db Database) GetReserveById(reserveId int) (models.Reserve, error) {
	reserve := models.Reserve{}

	query := `SELECT * FROM reserve WHERE id = $1;`
	row := db.Conn.QueryRow(query, reserveId)
	switch err := row.Scan(&reserve.ID, &reserve.IdBalance, &reserve.Summa, &reserve.DateReserve, &reserve.IdOrder, &reserve.IdService); err {
	case sql.ErrNoRows:
		return reserve, ErrNoMatch
	default:
		return reserve, err
	}
}

func (db Database) DeleteReserve(reserveId int) error {
	query := `DELETE FROM reserve WHERE id = $1;`
	_, err := db.Conn.Exec(query, reserveId)
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
