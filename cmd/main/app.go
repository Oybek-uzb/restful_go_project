package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"restful_go_project/internal/user"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("create router")
	router := httprouter.New()

	handler := user.NewHandler()
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatalln(server.Serve(listener))
}
