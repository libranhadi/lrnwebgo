package test

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"io"
	"lrnwebgo"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", lrnwebgo.Home)

	// test panic
	mux.HandleFunc("/panic", func (w http.ResponseWriter, req *http.Request)  {
		panic("TEST")
	})

	logMiddleware := &lrnwebgo.LogMiddleware {
		Handler : mux,
	}

	errorHandler := &lrnwebgo.HandleError {
		Handler : logMiddleware,
	}

	server := http.Server{
		Addr : "localhost:8000",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestMiddlewareHttpRouter(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		panic(err)
	}
	res := httptest.NewRecorder()

	router := httprouter.New()
	logMiddleware := &lrnwebgo.LogMiddleware {
		Handler : router,
	}
	router.GET("/", lrnwebgo.Index)

	logMiddleware.ServeHTTP(res, req);
	response := res.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "HELLO", string(body), "TEST")

}