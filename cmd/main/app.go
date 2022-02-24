package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"restful_go_project/internal/config"
	"restful_go_project/internal/user"
	"restful_go_project/pkg/logging"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logging.Init()
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("start application")

	if cfg.Listen.Type == "sock" {
		filepath.Abs(filepath.Dir(os.Args[0]))
	} else {

	}

	listener, err := net.Listen("tcp", cfg.Listen.Port)
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
