package main

import (
	"net/http"

	"github.com/SwanHtetAungPhyo/finance_track/stock_cal/logging"
	websocket_server "github.com/SwanHtetAungPhyo/finance_track/stock_cal/websocket"
)

func main(){
	logging.GlobalLogInit()
	http.HandleFunc("/ws",websocket_server.StockPriceHandler)
	logging.L().Info("Websocket server is starting")
	err := http.ListenAndServe(":7001",nil)
	if err != nil{
		logging.L().Fatal(err.Error())
	}	
}