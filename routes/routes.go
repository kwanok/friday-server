package routes

import (
	"fmt"
	"friday/endpoints/admin"
	"friday/tools"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v\n", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, msg)
	}
}

func Routes(r *gin.Engine) {
	err := godotenv.Load(".env")
	tools.ErrorHandler(err)

	r.GET("/", func(c *gin.Context) {
		c.String(200, "We got Gin")
	})

	rAdmin := r.Group("/admin")
	rAdmin.Use()
	{
		rAdmin.GET("/users", func(c *gin.Context) {
			admin.GetUsers(c)
		})
	}

	r.GET("/ws", func(c *gin.Context) {
		socketHandler(c.Writer, c.Request)
	})
}