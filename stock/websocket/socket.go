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
package swan_lib

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func (j *GlobalJWTMiddleware) FiberAuthorize() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Authorization Header"})
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing Bearer Token"})
		}

		tokenString := parts[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Method.Alg())
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				return nil, fmt.Errorf("could not extract claims")
			}


			if exp, ok := claims["exp"].(float64); ok {
				if time.Now().Unix() > int64(exp) {
					return nil, fmt.Errorf("token has expired")
				}
			}


			return []byte(j.Secret), nil
		})


		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or Expired Token"})
		}


		if !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Token"})
		}

		return c.Next()
	}
}

