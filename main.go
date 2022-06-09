package main

import (
	"net/http"

	"github.com/iqbalnzls/watchcommerce/src/delivery"
	inHttp "github.com/iqbalnzls/watchcommerce/src/delivery/http"
)

func main() {
	mux := http.NewServeMux()

	inHttp.StartHttpService(mux, delivery.SetupContainer())
}
