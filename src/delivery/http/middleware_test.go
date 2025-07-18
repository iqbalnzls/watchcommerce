package http_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	inHttp "github.com/iqbalnzls/watchcommerce/src/delivery/http"
	"github.com/iqbalnzls/watchcommerce/src/shared/constant"
)

func TestSetupMiddleware(t *testing.T) {
	tests := []struct {
		name           string
		authorization  string
		expectedCode   int
		expectedHeader http.Header
	}{
		{
			name:          "Success - Valid request",
			authorization: "",
			expectedCode:  http.StatusOK,
			expectedHeader: http.Header{
				"Access-Control-Allow-Origin":  []string{"*"},
				"Access-Control-Allow-Headers": []string{"Content-Type"},
				"Content-Type":                 []string{"application/json"},
			},
		},
		{
			name:          "Success - Valid swagger auth",
			authorization: "asdjkhNasdb90834aSD",
			expectedCode:  http.StatusOK,
			expectedHeader: http.Header{
				"Access-Control-Allow-Origin":  []string{"*"},
				"Access-Control-Allow-Headers": []string{"Content-Type"},
				"Content-Type":                 []string{"application/json"},
			},
		},
		{
			name:           "Error - Invalid swagger auth",
			authorization:  "invalid-auth",
			expectedCode:   http.StatusForbidden,
			expectedHeader: http.Header{},
		},
	}

	c := context.Background()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test handler that checks the context
			testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Verify context is set
				ctx := r.Context()
				appCtx := ctx.Value(constant.AppContext)
				assert.NotNil(t, appCtx, "AppContext should be set")

				w.WriteHeader(http.StatusOK)
			})

			// Setup the middleware
			middleware := inHttp.SetupMiddleware(c)
			handler := middleware(testHandler)

			// Create test request
			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			if tt.authorization != "" {
				req.Header.Set("Authorization-Swagger", tt.authorization)
			}

			rr := httptest.NewRecorder()

			// Serve the request
			handler.ServeHTTP(rr, req)

			// Assert response code
			assert.Equal(t, tt.expectedCode, rr.Code)

			// Assert headers for successful requests
			if tt.expectedCode == http.StatusOK {
				for key, value := range tt.expectedHeader {
					assert.Equal(t, value[0], rr.Header().Get(key))
				}
			}
		})
	}
}
