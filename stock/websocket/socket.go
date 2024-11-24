package websocket_server

import (
	"net/http"
	"math/rand"
	"time"

	"github.com/SwanHtetAungPhyo/finance_track/stock_cal/logging"
	"github.com/gorilla/websocket"
)

const (
	stockSymbol     = "AAPL"
	initialPrice    = 150.00
	priceFluctuation = 2.0
	updateInterval  = 1 *time.Second
	redisAddress    = "localhost:6379"
)


var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func generateSimulation(currentPrice float64) float64 {
	priceChange := rand.Float64()*priceFluctuation*2 - priceFluctuation
	currentPrice = currentPrice + priceChange
	if currentPrice < 0 {
		currentPrice = 0 
	}
	return currentPrice
}

func StockPriceHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logging.L().Error(err.Error())
		return
	}

	defer conn.Close()

	currentPrice := initialPrice
	initialData := map[string]interface{}{
		"symbol": stockSymbol,
		"price":  currentPrice,
	}
	err = conn.WriteJSON(initialData)
	if err != nil {
		logging.L().Error(err.Error())
		return 
	}

	for {
		currentPrice = generateSimulation(currentPrice)

		updateData := map[string]interface{}{
			"symbol": stockSymbol,
			"price":  currentPrice,
		}
		err = conn.WriteJSON(updateData)
		if err != nil {
			logging.L().Error( err.Error())
			return
		}
		time.Sleep(updateInterval)
	}

}
