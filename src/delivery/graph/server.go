package graph

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"

	"github.com/iqbalnzls/watchcommerce/src/delivery"
)

func StartGraphQLServer(ctx context.Context, container *delivery.Container) {
	mux := http.NewServeMux()

	srv := handler.New(NewExecutableSchema(
		Config{
			Resolvers: SetupResolver(container.ProductService, container.BrandService, container.Validator),
		}),
	)

	// Register transports explicitly
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})

	mux.Handle("/", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", SetupMiddleware(srv))

	server := http.Server{
		Addr:    container.Config.Apps.GetGraphQLAddress(),
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("GraphQL Server Started on Port : ", container.Config.Apps.GraphQLPort)

	<-ctx.Done()
	log.Print("GraphQL Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}

	log.Print("GraphQL Server Exited Properly")
}
