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

var balanceIDKey = "balanceID"

func balances(router chi.Router) {
	router.Get("/", getAllBalances)
	router.Post("/", createBalance)

	router.Route("/{balanceId}", func(router chi.Router) {
		router.Use(BalanceContext)
		router.Get("/", getBalance)
		router.Put("/", updateBalance)
		router.Delete("/", deleteBalance)
	})
}

func BalanceContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		balanceId := chi.URLParam(r, "balanceId")
		if balanceId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("balance ID is required")))
			return
		}
		id, err := strconv.Atoi(balanceId)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid balance ID")))
		}
		ctx := context.WithValue(r.Context(), balanceIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllBalances(w http.ResponseWriter, r *http.Request) {
	balances, err := dbInstance.GetAllBalances()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, balances); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

func createBalance(w http.ResponseWriter, r *http.Request) {
	balance := &models.Balance{}
	//log.Printf("kek")
	/*if err := render.Bind(r, balance); err != nil {
		log.Printf("%d", &balance.IdUser)
		render.Render(w, r, ErrBadRequest)
		return
	}*/

	//log.Printf("%d", &balance.)

	if err := dbInstance.AddBalance(balance); err != nil {



		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, balance); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func getBalance(w http.ResponseWriter, r *http.Request) {
	balanceID := r.Context().Value(balanceIDKey).(int)
	balance, err := dbInstance.GetBalanceById(balanceID)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &balance); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func deleteBalance(w http.ResponseWriter, r *http.Request) {
	balanceId := r.Context().Value(balanceIDKey).(int)
	err := dbInstance.DeleteBalance(balanceId)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
}

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
