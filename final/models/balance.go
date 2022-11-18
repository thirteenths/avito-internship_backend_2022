package models

import (
	//"fmt"
	"net/http"
)

type Balance struct{
	ID int `json:"id"`
	IdUser int `json:"id_user"`
	Summa float64 `json:"summa"`
	DateUpgrate string `json:"date_upgrate"`
}

type BalanceList struct {
	Balances []Balance `json:"balances"`
}

func (i *Balance) Bind(r *http.Request) error {
	//if i.Summa == "" {
	//	return fmt.Errorf("name is a required field")
	//}
	return nil
}

func (*BalanceList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Balance) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
