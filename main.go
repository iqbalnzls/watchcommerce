package main

import (
	"net/http"

	_ "github.com/iqbalnzls/watchcommerce/docs"
	"github.com/iqbalnzls/watchcommerce/src/delivery"
	inGraphQL "github.com/iqbalnzls/watchcommerce/src/delivery/graph"
	inHttp "github.com/iqbalnzls/watchcommerce/src/delivery/http"
)

// @title Watchcommerce API Documentation
// @version 1.0
// @description This is api documentation for watchcommerce.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @BasePath /
func main() {
	mux := http.NewServeMux()

	container := delivery.SetupContainer()

	go inGraphQL.StartGraphQLService(mux, container)

	inHttp.StartHttpService(mux, container)
}
