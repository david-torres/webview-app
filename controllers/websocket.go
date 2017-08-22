package controllers

import (
	"io"
	"log"

	"golang.org/x/net/websocket"
)

// Message to receive/send
type Message struct {
	Message string `json:"msg"`
}

// Socket handles the websocket
func Socket() websocket.Handler {
	return websocket.Handler(func(ws *websocket.Conn) {
		// loop forever receiving from socket
		for {
			incoming := new(Message)
			err := websocket.JSON.Receive(ws, incoming)
			if err != nil {
				if err == io.EOF {
					log.Println("websocket received EOF")
					continue
				}

				log.Println("websocket receive error")
				log.Println(err.Error())
				continue
			}

			log.Println("Incoming message: " + incoming.Message)

			outgoing := new(Message)
			outgoing.Message = "pong"
			err = websocket.JSON.Send(ws, outgoing)

			if err != nil {
				log.Println("websocket send error")
				log.Println(err.Error())
				continue
			}

			log.Println("Outgoing message: " + outgoing.Message)
		}
	})
}
