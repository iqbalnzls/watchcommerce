package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/mbndr/figlet4go"

	_ "github.com/iqbalnzls/watchcommerce/docs"
	"github.com/iqbalnzls/watchcommerce/src/delivery"
	inGraphQL "github.com/iqbalnzls/watchcommerce/src/delivery/graph"
	inHttp "github.com/iqbalnzls/watchcommerce/src/delivery/http"
	"github.com/iqbalnzls/watchcommerce/src/shared/constant"
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
	// Create a context that is cancelled on interrupt
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Channel to listen for interrupt signals
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)

	// Init dependencies
	container := delivery.SetupContainer()

	fmt.Print(showBanner())
	fmt.Printf(constant.AppVersion + "\n\n")

	// Start graphql and http servers
	go inGraphQL.StartGraphQLServer(ctx, container)
	go inHttp.StartHttpServer(ctx, container)

	// Block until an interrupt signal is received
	<-signalChan

	// Cancel the context (notify all goroutines to stop)
	cancel()

}

func showBanner() string {
	ascii := figlet4go.NewAsciiRender()

	options := figlet4go.NewRenderOptions()
	options.FontName = "larry3d"
	options.FontColor = []figlet4go.Color{
		figlet4go.ColorRed,
		figlet4go.ColorBlue,
		figlet4go.ColorCyan,
		figlet4go.ColorMagenta,
	}

	renderStr, _ := ascii.RenderOpts("watchcommerce", options)

	return renderStr
}
