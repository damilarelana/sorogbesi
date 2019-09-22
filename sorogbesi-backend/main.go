package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"sorogbesi-backend/pkg/websocket"

	"github.com/gorilla/mux"
)

// define server http homePageEndpoint
func homePageEndpointHandler(w http.ResponseWriter, r *http.Request) {
	dataToRender := "Websocket Server: Home Page"
	io.WriteString(w, dataToRender)
}

func wsEndpointHandlerWithPool(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {

	wsConn, err := websocket.Upgrade(w, r) //convert the http to ws i.e. convert the w to wsConn
	if err != nil {
		panic(err.Error())
	}
	log.Println("Server: Client successfully connected via WebSockets ... ")

	// create a client for every connection
	client := &websocket.Client{
		Conn: wsConn,
		Pool: pool,
	}

	// register the client
	pool.Register <- client
	client.Read()
}

func serviceRequestHandlers() {

	// initialize a new websocket Pool
	pool := websocket.NewPool()
	go pool.Start()

	newRouter := mux.NewRouter().StrictSlash(true)
	newRouter.HandleFunc("/", homePageEndpointHandler)
	newRouter.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		wsEndpointHandlerWithPool(pool, w, r)
	})
	log.Fatal(http.ListenAndServe(":8080", newRouter))
}

func main() {
	fmt.Println("Server: Sorogbesi backend server initialising")
	serviceRequestHandlers()
}
