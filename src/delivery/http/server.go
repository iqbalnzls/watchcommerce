package http

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/iqbalnzls/watchcommerce/src/delivery"
)

func StartHttpServer(ctx context.Context, container *delivery.Container) {
	mux := http.NewServeMux()

	SetupRouter(mux, SetupMiddleware(ctx), SetupHandler(container))

	srv := http.Server{
		Addr:    container.Config.Apps.GetHttpAddress(),
		Handler: mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Http Server Started on Port : ", container.Config.Apps.HttpPort)

	<-ctx.Done()
	log.Print("Http Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	log.Print("Http Server Exited Properly")

}
