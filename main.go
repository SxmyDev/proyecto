package main

import (
	"github.com/Sxmmy2030/proyecto/auth"
	"github.com/Sxmmy2030/proyecto/websocket"
	"net/http"
)

func main() {
	http.HandleFunc("/login", auth.LoginHandler)
	http.HandleFunc("/register", auth.RegisterHandler)
	http.HandleFunc("/ws", websocket.WsHandler)

	http.ListenAndServe(":8080", nil)
}
