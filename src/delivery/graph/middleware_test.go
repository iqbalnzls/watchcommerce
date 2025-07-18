package graph_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/iqbalnzls/watchcommerce/src/delivery/graph"
	"github.com/iqbalnzls/watchcommerce/src/shared/constant"
)

func TestSetupMiddleware(t *testing.T) {
	tests := []struct {
		name         string
		path         string
		expectedPath string
		expectedCode int
	}{
		{
			name:         "Success - Middleware adds context",
			path:         "/graphql",
			expectedPath: "/graphql",
			expectedCode: http.StatusOK,
		},
		{
			name:         "Success - Different path",
			path:         "/api/products",
			expectedPath: "/api/products",
			expectedCode: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test handler that checks the context
			testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Verify context is set
				ctx := r.Context()
				appCtx := ctx.Value(constant.AppContext)
				assert.NotNil(t, appCtx, "AppContext should be set")

				// Verify request path
				assert.Equal(t, tt.expectedPath, r.URL.Path)

				w.WriteHeader(tt.expectedCode)
			})

			// Setup the middleware
			handler := graph.SetupMiddleware(testHandler)

			// Create test request
			req := httptest.NewRequest(http.MethodPost, tt.path, nil)
			rr := httptest.NewRecorder()

			// Serve the request
			handler.ServeHTTP(rr, req)

			// Assert response
			assert.Equal(t, tt.expectedCode, rr.Code)
		})
	}
}
