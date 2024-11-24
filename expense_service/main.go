package main

import (
	"net/http"

	"github.com/SwanHtetAungPhyo/expense_service/database"
	"github.com/SwanHtetAungPhyo/expense_service/handler"
)


func main(){
	database.DB_INIT()
	database.Migration(&handler.Expense{})
	
	Http_Server()
}
func Http_Server(){
	mux := http.NewServeMux()
	handlers := handler.NewHttpHandler()
	mux.HandleFunc("/expense",handlers.CreateExpense) 
	http.ListenAndServe(":5002",mux)
}


