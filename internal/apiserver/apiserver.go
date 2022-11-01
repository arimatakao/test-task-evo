package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type APIServer struct {
	db     *sql.DB
	router *mux.Router
}

func (srv *APIServer) SetDB(db *sql.DB) {
	srv.db = db
}

func (srv *APIServer) Start() error {
	srv.router = mux.NewRouter()
	srv.router.HandleFunc("/", srv.handleHello())
	srv.router.HandleFunc("/transactionid/{id:[0-9]+}", srv.handleGetByTransactionId()).Methods("GET")
	srv.router.HandleFunc("/terminalid", srv.handleGetByTerminalId()).Methods("GET")
	srv.router.HandleFunc("/status", srv.handleGetByStatus()).Methods("GET")
	srv.router.HandleFunc("/paymenttype", srv.handleGetByPaymentType()).Methods("GET")
	srv.router.HandleFunc("/date", srv.handleGetByDate()).Methods("GET")
	srv.router.HandleFunc("/paymentnarrative", srv.handleGetByPaymentNarrative()).Methods("GET")

	return http.ListenAndServe(":8080", srv.router)
}
