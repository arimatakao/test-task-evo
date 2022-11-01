package apiserver

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/arimatakao/test-task-evo/internal/model"
	"github.com/gorilla/mux"
)

func (srv *APIServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Test task for EVO Finctech")
	}
}

// return 1 row
func (srv *APIServer) handleGetByTransactionId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		vars := mux.Vars(r)
		if _, ok := vars["id"]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{ "error": "Bad id" }`)
		}
		resp := model.Info{}
		rows, err := srv.db.Query("SELECT * FROM meterial WHERE TransactionId=$1", vars["id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, `{ "error": "Internal error" }`)
		}
		defer rows.Close()
		// error is here
		for rows.Next() {
			err = rows.Scan(resp.TransactionId)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				io.WriteString(w, `{ "error": "Bad bind struct" }`)
				return
			}
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
	}
}

// return many rows
func (srv *APIServer) handleGetByTerminalId() http.HandlerFunc {
	type request struct {
		TerminalId []string `json:"TerminalId"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		req := request{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{ "error": "Wrong json" }`)
		}
		io.WriteString(w, "not implemented")
	}
}

func (srv *APIServer) handleGetByStatus() http.HandlerFunc {
	type request struct {
		Status string `json:"Status"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		req := request{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{ "error": "Wrong json" }`)
		}
		io.WriteString(w, "not implemented")
	}
}

func (srv *APIServer) handleGetByPaymentType() http.HandlerFunc {
	type request struct {
		PaymentType string `json:"PaymentType"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		req := request{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{ "error": "Wrong json" }`)
		}
		io.WriteString(w, "not implemented")
	}
}

func (srv *APIServer) handleGetByDate() http.HandlerFunc {
	type request struct {
		DatePost string `json:"DatePost"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		req := request{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{ "error": "Wrong json" }`)
		}
		io.WriteString(w, "not implemented")
	}
}

func (srv *APIServer) handleGetByPaymentNarrative() http.HandlerFunc {
	type request struct {
		PaymentNarrative string `json:"PaymentNarrative"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		req := request{}
		json.NewDecoder(r.Body).Decode(&req)
		io.WriteString(w, "not implemented")
	}
}
