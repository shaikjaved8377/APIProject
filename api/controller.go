package api

import (
	"Project/model"
	"database/sql"
	"encoding/json"
	"net/http"
)

type Handler struct {
	biz IbizLogic
}

func NewHandler(db *sql.DB) Handler {
	return Handler{biz: NewBizLogic(db)}

}
func (h Handler) Createhandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return

		}

		var book model.Book
		if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := h.biz.CreateBookLogic(book); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}
		w.WriteHeader(http.StatusOK)
	}
}
