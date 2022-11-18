package handler

import (
	//"log"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/thirteenths/final/db"
	"github.com/thirteenths/final/models"
)

var PayrollIDKey = "payrollID"

func payrolles(router chi.Router) {
	router.Get("/", getAllBalances)
	router.Post("/", createBalance)

	router.Route("/{payrollID}", func(router chi.Router) {
		router.Use(BalanceContext)
		router.Get("/", getBalance)
		router.Put("/", updateBalance)
		router.Delete("/", deleteBalance)
	})
}

func PayrollContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		payrollId := chi.URLParam(r, "payrollID")
		if payrollId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("payroll ID is required")))
			return
		}
		id, err := strconv.Atoi(payrollId)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid payroll ID")))
		}
		ctx := context.WithValue(r.Context(), PayrollIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllPayrolles(w http.ResponseWriter, r *http.Request) {
	payrolles, err := dbInstance.GetAllPayrolles()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, payrolles); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

func createPayroll(w http.ResponseWriter, r *http.Request) {
	payroll := &models.Payroll{}
	//log.Printf("kek")
	/*if err := render.Bind(r, balance); err != nil {
		log.Printf("%d", &balance.IdUser)
		render.Render(w, r, ErrBadRequest)
		return
	}*/

	//balance := &models.Balance{}


	if err := dbInstance.AddPayroll(payroll); err != nil {

		//log.Panicf("%d", &payroll.IdBalance)

		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, payroll); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func getPayroll(w http.ResponseWriter, r *http.Request) {
	payrollID := r.Context().Value(PayrollIDKey).(int)
	payroll, err := dbInstance.GetPayrollById(payrollID)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &payroll); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func deletePayroll(w http.ResponseWriter, r *http.Request) {
	payrollID := r.Context().Value(PayrollIDKey).(int)
	err := dbInstance.DeletePayroll(payrollID)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
}
/*
func updateBalance(w http.ResponseWriter, r *http.Request) {
	balanceId := r.Context().Value(balanceIDKey).(int)
	balanceData := models.Balance{}
	if err := render.Bind(r, &balanceData); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	balance, err := dbInstance.UpdateBalance(balanceId, balanceData)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &balance); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
*/