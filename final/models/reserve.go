package models

import (
	//"fmt"
	"net/http"
)

type Reserve struct{
	ID int `json:"id"`
	IdBalance int `json:"id_balance"`
	Summa float64 `json:"summa"`
	DateReserve string `json:"date_reserve"`
	IdOrder int `json:"id_order"`
	IdService int `json:"id_service"`
}

type ReserveList struct {
	Reserves []Reserve `json:reserves"`
}

func (i *Reserve) Bind(r *http.Request) error {
	//if i.Summa == "" {
	//	return fmt.Errorf("name is a required field")
	//}
	return nil
}

func (*ReserveList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Reserve) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
