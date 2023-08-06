package test

import(
	"testing"
	"net/http"
	"net/http/httptest"
	"io"
	"lrnwebgo/resources"
	"github.com/stretchr/testify/assert"
	"github.com/julienschmidt/httprouter"
)


func TestServeFile(t *testing.T)  {
	a := &resources.Resources

	router := httprouter.New()

	router.ServeFiles("/files/*filepath", http.FS(a))

	req, err := http.NewRequest("GET", "/files/test.txt", nil)
	if err != nil {
		panic(err)
	}
	res := httptest.NewRecorder()

	router.ServeHTTP(res, req);
	
	response := res.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "WELCOME", string(body), "TEST")
	assert.Equal(t, 200, response.StatusCode, "OKEE")	
}


