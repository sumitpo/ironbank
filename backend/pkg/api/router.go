package api

import (
	"bank-app/pkg/database"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(db *database.MySQLDB) *mux.Router {
	r := mux.NewRouter()

	// Auth routes
	r.HandleFunc("/auth/login", loginHandler(db)).Methods("POST")

	// Account routes
	r.HandleFunc("/accounts", authMiddleware(getAccountsHandler(db))).Methods("GET")
	r.HandleFunc("/accounts/{id}", authMiddleware(getAccountHandler(db))).Methods("GET")

	// Transaction routes
	r.HandleFunc("/transactions", authMiddleware(createTransactionHandler(db))).Methods("POST")

	return r
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if !isValidToken(token) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}
