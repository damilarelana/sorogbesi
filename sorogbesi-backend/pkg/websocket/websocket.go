package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// define an upgrader i.e. attributes of the upgrader used to convert the http.conn to websocket.conn
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Upgrade is the functin the server uses to turn a normal http into a websocket connection
func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {

	// required to handle CORS i.e. check if request from a different domain is allowed to connect
	upgrader.CheckOrigin = func(r *http.Request) bool { // the anonymous function returns a true or false i.e. true being allowed
		return true //let's assume for testing purposes that the requesting domain is allowed, reqgardless of origin/source
	}

	wsConn, err := upgrader.Upgrade(w, r, nil) //convert the http to ws i.e. convert the w to wsConn
	if err != nil {
		log.Println(err)
		return wsConn, err
	}
	return wsConn, nil
}
