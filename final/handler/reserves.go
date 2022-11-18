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

var ReserveIDKey = "reserveID"

func reserves(router chi.Router) {
	router.Get("/", getAllBalances)
	router.Post("/", createBalance)

	router.Route("/{reserveID}", func(router chi.Router) {
		router.Use(BalanceContext)
		router.Get("/", getBalance)
		router.Put("/", updateBalance)
		router.Delete("/", deleteBalance)
	})
}

func ReserveContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reserveId := chi.URLParam(r, "reserveId")
		if reserveId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("reserve ID is required")))
			return
		}
		id, err := strconv.Atoi(reserveId)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid reserve ID")))
		}
		ctx := context.WithValue(r.Context(), ReserveIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllReserves(w http.ResponseWriter, r *http.Request) {
	reserves, err := dbInstance.GetAllReserves()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, reserves); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

func createReserve(w http.ResponseWriter, r *http.Request) {
	reserve := &models.Reserve{}
	//log.Printf("kek")
	/*if err := render.Bind(r, balance); err != nil {
		log.Printf("%d", &balance.IdUser)
		render.Render(w, r, ErrBadRequest)
		return
	}*/
	if err := dbInstance.AddReserve(reserve); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, reserve); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func getReserve(w http.ResponseWriter, r *http.Request) {
	reserveID := r.Context().Value(ReserveIDKey).(int)
	reserve, err := dbInstance.GetReserveById(reserveID)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &reserve); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func deleteReserve(w http.ResponseWriter, r *http.Request) {
	reserveID := r.Context().Value(ReserveIDKey).(int)
	err := dbInstance.DeleteReserve(reserveID)
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