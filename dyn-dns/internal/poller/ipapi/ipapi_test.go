package ipapi

import (
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/larwef/my-platform/dyn-dns/internal/poller"
	"github.com/stretchr/testify/assert"
)

func TestIPApi_Poll(t *testing.T) {
	tests := []struct {
		name      string
		handler   http.Handler
		expect    net.IP
		expectErr error
	}{
		{
			name: "Test error from api",
			handler: http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
				http.Error(rw, "some error", http.StatusInternalServerError)
			}),
			expectErr: &poller.Error{Code: http.StatusInternalServerError, Message: "some error\n"},
		},
		{
			name: "Test http error 429 from api",
			handler: http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
				http.Error(rw, "too many requests", http.StatusTooManyRequests)
			}),
			expectErr: &poller.Error{Code: http.StatusTooManyRequests, Message: "too many requests\n"},
		},
		{
			name: "Test successful",
			handler: http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
				f, err := os.Open("../../../test/data/response.json")
				assert.NoError(t, err)
				_, _ = io.Copy(rw, f)
			}),
			expect: net.ParseIP("85.166.235.188"),
		},
	}
	for _, test := range tests {
		srv := httptest.NewServer(test.handler)
		ipAPI := &IPApi{url: srv.URL, client: &http.Client{}}
		res, err := ipAPI.Poll()
		assert.Equal(t, test.expect, res)
		assert.Equal(t, test.expectErr, err)
	}
}
