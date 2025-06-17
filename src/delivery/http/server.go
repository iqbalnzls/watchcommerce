package http

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/iqbalnzls/watchcommerce/src/delivery"
)

func StartHttpServer(mux *http.ServeMux, container *delivery.Container) {
	SetupRouter(mux, SetupMiddleware(), SetupHandler(container))

	srv := http.Server{
		Addr:    container.Config.Apps.GetHttpAddress(),
		Handler: mux,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Http Server Started on Port : ", container.Config.Apps.HttpPort)

	<-done
	log.Print("Http Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	log.Print("Http Server Exited Properly")

}
