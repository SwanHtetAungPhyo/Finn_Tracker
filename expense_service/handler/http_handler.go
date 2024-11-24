package handler

import (
	"encoding/json"
	"net/http"

	"github.com/SwanHtetAungPhyo/expense_service/grpc"
)
type Expense struct{
	ID int `json:"id"`
	UserId int `json:"user_id"`
	Amount int 	`json:"amount"`
}


type HttpHandler struct{
}


func NewHttpHandler() *HttpHandler{
	return &HttpHandler{}
}

func (h *HttpHandler) CreateExpense(w http.ResponseWriter, r *http.Request) {
	var expense Expense
	err := json.NewDecoder(r.Body).Decode(&expense)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}


	exists, err := grpc.CheckUserExist(uint(expense.UserId))
	if err != nil || !exists {
		http.Error(w, "User does not exist", http.StatusNotFound)
		return
	}
	expense.ID = 1 // This would normally come from the database

	responseData, err := json.Marshal(expense)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseData)
}