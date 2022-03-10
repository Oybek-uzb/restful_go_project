package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	author2 "restful_go_project/internal/author"
	author "restful_go_project/internal/author/db"
	"restful_go_project/internal/config"
	"restful_go_project/internal/user"
	"restful_go_project/pkg/client/postgresql"
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

	postgreSQLClient, err := postgresql.NewClient(context.TODO(), 3, cfg.Storage)
	if err != nil {
		logger.Fatalf("%s", err)
	}

	repository := author.NewRepository(postgreSQLClient, logger)

	newAu := author2.Author{
		Name: "Tohir Malik",
	}
	err = repository.Create(context.TODO(), &newAu)
	if err != nil {
		logger.Fatalf("%v", err)
	}
	logger.Infof("%v", newAu)

	foundAu, err := repository.FindOne(context.TODO(), "89f44248-c821-40d6-bb1c-9fc7d1f747c4")
	if err != nil {
		logger.Fatalf("%v", err)
	}
	logger.Infof("%v", foundAu)

	all, err := repository.FindAll(context.TODO())
	if err != nil {
		logger.Fatalf("%v", err)
	}

	for _, au := range all {
		logger.Infof("%v", au)
	}

	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("start application")

	var listener net.Listener
	var listenErr error

	logger.Info("detect app path")
	if cfg.Listen.Type == "sock" {
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("create socket")
		socketPath := path.Join(appDir, "app:sock")

		logger.Info("listen unix socket")
		listener, listenErr = net.Listen("unix", socketPath)
		logger.Infof("server is listening unix socket: %s", socketPath)
	} else {
		logger.Info("listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		logger.Infof("server is listening %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	logger.Fatal(server.Serve(listener))
}
