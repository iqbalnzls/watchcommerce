package main

import (
	"net/http"

	"github.com/iqbalnzls/watchcommerce/src/delivery"
	inHttp "github.com/iqbalnzls/watchcommerce/src/delivery/http"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	mux := http.NewServeMux()

	inHttp.StartHttpService(mux, delivery.SetupContainer())
}
