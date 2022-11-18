package models

import (
	//"fmt"
	"net/http"
)

type Payroll struct{
	ID int `json:"id"`
	IdBalance int `json:"id_balance"`
	Summa float64 `json:'summa"`
	DatePayroll string `json:"date_payroll"`
}

type PayrollList struct{
	Payrolles []Payroll `json:"payrolles`
}

func (i *Payroll) Bind(r *http.Request) error {
	//if i.Summa == "" {
	//	return fmt.Errorf("name is a required field")
	//}
	return nil
}

func (*PayrollList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Payroll) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
