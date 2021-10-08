package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"
	"time"
)

var upgrader = websocket.Upgrader{}

func main() {
	port := ":" + os.Getenv("PORT")

	e := echo.New()
	e.Use(middleware.Logger())
	e.Pre(middleware.RemoveTrailingSlash())

	e.Logger.SetLevel(log.DEBUG)

	e.GET("/ws", geoWebsocket)

	e.Logger.Fatal(e.Start(port))
}

func geoWebsocket(c echo.Context) error {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		_, message, err := ws.ReadMessage()
		if err != nil {
			panic(err)
		}
		fmt.Println(len(string(message)))
		time.Sleep(1 * time.Second)
	}
	return nil
}
