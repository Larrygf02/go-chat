package main

import (
	"fmt"
	"log"
	"net/http"

	engineio "github.com/googollee/go-engine.io"
	"github.com/googollee/go-engine.io/transport"
	"github.com/googollee/go-engine.io/transport/polling"
	"github.com/googollee/go-engine.io/transport/websocket"
	socketio "github.com/googollee/go-socket.io"
	"github.com/rs/cors"
)

func main() {
	fmt.Println("Hello chat")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
	})

	pt := polling.Default

	wt := websocket.Default
	wt.CheckOrigin = func(req *http.Request) bool {
		return true
	}
	server, err := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			pt,
			wt,
		},
	})

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
	mux.Handle("/socket.io/", server)
	//http.Handle("/", http.FileServer(http.Dir("./public")))
	log.Println("Server on Port 3000")
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowedMethods:   []string{"GET", "PUT", "OPTIONS", "POST", "DELETE"},
		AllowCredentials: true,
	})
	// decorate existing handler with cors functionality set in c
	handler := c.Handler(mux)
	log.Fatal(http.ListenAndServe(":3000", handler))
	//log.Fatal(http.ListenAndServe(":3000", nil))
}
