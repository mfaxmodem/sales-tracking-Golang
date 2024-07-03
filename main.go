package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"log"
	"net/http"
	"sales-tracking/config"
	"sales-tracking/models"
	"sales-tracking/routes"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to set websocket upgrade:", err)
		return
	}
	defer func(ws *websocket.Conn) {
		err := ws.Close()
		if err != nil {

		}
	}(ws)

	db, exists := c.Get("db")
	if !exists {
		log.Println("Database connection not found in context")
		return
	}

	for {
		var msg map[string]interface{}
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Println("Error reading JSON:", err)
			break
		}
		log.Printf("Received message: %v", msg)

		// Save received data to database
		if msg["type"] == "checkpoint" {
			var checkpoint models.Checkpoint
			if err := mapToStruct(msg["data"], &checkpoint); err != nil {
				log.Println("Error converting map to struct:", err)
				continue
			}
			db.(*gorm.DB).Create(&checkpoint)
		}
		// Add similar blocks for other models (Route, Store, User, Visit)

		err = ws.WriteJSON(msg)
		if err != nil {
			log.Println("Error writing JSON:", err)
			break
		}
	}
}

func mapToStruct(data interface{}, result interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, result)
}

func main() {
	// اتصال به پایگاه داده
	config.ConnectDatabase()

	// مهاجرت مدل‌ها
	err := config.DB.AutoMigrate(&models.User{}, &models.Route{}, &models.Checkpoint{}, &models.Store{}, &models.Visit{})
	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	router := gin.Default()

	// اضافه کردن اتصال پایگاه داده به کانتکست Gin
	router.Use(func(c *gin.Context) {
		c.Set("db", config.DB)
		c.Next()
	})

	// تنظیم مسیرهای HTTP
	routes.SetupRoutes(router)

	// تنظیم مسیر WebSocket
	router.GET("/ws", handleConnections)

	err = router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
