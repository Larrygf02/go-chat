package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	fmt.Println("Hello chat")
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	// sockets
	server.OnConnect("/", func(so socketio.Conn) error {
		so.SetContext("")
		so.Join("chat")
		fmt.Println("Connected: ", so.ID())
		return nil
	})

	server.OnEvent("/", "chat message", func(s socketio.Conn, msg string) {
		fmt.Println("messsage:", msg)
		// Emite al usuario mismo
		// s.Emit("chat message", msg)
		// Emite el mensaje a todos los usuarios de la sala
		server.BroadcastToRoom("", "chat", "chat message", msg)
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	go server.Serve()
	defer server.Close()
	//http
	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Server on Port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
