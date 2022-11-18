package db

import (
	"database/sql"

	"github.com/thirteenths/final/models"
)

func (db Database) GetAllBalances() (*models.BalanceList, error) {
	list := &models.BalanceList{}

	rows, err := db.Conn.Query("SELECT * FROM balance")
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var balance models.Balance
		err := rows.Scan(&balance.ID, &balance.IdUser, &balance.Summa, &balance.DateUpgrate)
		if err != nil {
			return list, err
		}
		list.Balances = append(list.Balances, balance)
	}
	return list, nil
}

func (db Database) AddBalance(balance *models.Balance) error {
	var id int
	var dataUpgrate string
	query := `INSERT INTO balance (id_user, summa) VALUES ($1, $2) RETURNING id, date_upgrate`
	err := db.Conn.QueryRow(query, balance.IdUser, balance.Summa).Scan(&id, &dataUpgrate)
	if err != nil {
		return err
	}

	balance.ID = id
	balance.DateUpgrate = dataUpgrate
	return nil
}

func (db Database) GetBalanceById(balanceId int) (models.Balance, error) {
	balance := models.Balance{}

	query := `SELECT * FROM balance WHERE id = $1;`
	row := db.Conn.QueryRow(query, balanceId)
	switch err := row.Scan(&balance.ID, &balance.IdUser, &balance.Summa, &balance.DateUpgrate); err {
	case sql.ErrNoRows:
		return balance, ErrNoMatch
	default:
		return balance, err
	}
}

func (db Database) DeleteBalance(balanceId int) error {
	query := `DELETE FROM balance WHERE id = $1;`
	_, err := db.Conn.Exec(query, balanceId)
	switch err {
	case sql.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}

func (db Database) UpdateBalance(balanceId int, balanceData models.Balance) (models.Balance, error) {
	balance := models.Balance{}
	query := `UPDATE balance SET summa=$1 WHERE id=$2 RETURNING id, id_user, summa, data_upgrate`
	err := db.Conn.QueryRow(query, balanceData.IdUser, balanceData.Summa, balanceId).Scan(&balance.ID, &balance.IdUser, &balance.Summa, &balance.DateUpgrate)
	if err != nil {
		if err == sql.ErrNoRows {
			return balance, ErrNoMatch
		}
		return balance, err
	}

	return balance, nil
}
