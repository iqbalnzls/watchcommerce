package graph

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/iqbalnzls/watchcommerce/src/delivery"
)

func StartGraphQLService(mux *http.ServeMux, container *delivery.Container) {
	srv := handler.NewDefaultServer(NewExecutableSchema(Config{Resolvers: &Resolver{
		productService: container.ProductService,
		v:              container.Validator,
	}}))

	mux.HandleFunc("/", playground.Handler("GraphQL playground", "/query"))
	mux.HandleFunc("/query", srv.ServeHTTP)

	go func() {
		log.Fatal(http.ListenAndServe(container.Config.Apps.GetGraphQLAddress(), mux))
	}()

	log.Print("GraphQL Server Started on Port : ", container.Config.Apps.GraphQLPort)
}
