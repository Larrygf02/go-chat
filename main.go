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
		fmt.Println("Connected: ", so.ID())
		return nil
	})

	//http
	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Server on Port 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
