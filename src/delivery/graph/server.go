package graph

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/iqbalnzls/watchcommerce/src/delivery"
)

func StartGraphQLServer(mux *http.ServeMux, container *delivery.Container) {
	srv := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: &Resolver{
		productService: container.ProductService,
		v:              container.Validator,
	}}))

	mux.HandleFunc("/", playground.Handler("GraphQL playground", "/query"))
	mux.HandleFunc("/query", srv.ServeHTTP)

	server := http.Server{
		Addr:    container.Config.Apps.GetGraphQLAddress(),
		Handler: srv,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("GraphQL Server Started on Port : ", container.Config.Apps.GraphQLPort)

	<-done
	log.Print("GraphQL Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	log.Print("GraphQL Server Exited Properly")
}
