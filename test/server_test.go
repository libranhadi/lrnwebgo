package test

import(
	"testing"
	"net/http"
	"net/http/httptest"
	"lrnwebgo"
	"io"
	"github.com/stretchr/testify/assert"
)

func TestRunServe(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8000", nil)
	record := httptest.NewRecorder()
	lrnwebgo.Welcome(record, req)

	resp := record.Result()
	body, _ := io.ReadAll(resp.Body);

	assert.Equal(t, "WELCOME", string(body), "TEST")
	assert.Equal(t, 200, resp.StatusCode, "TEST")
}

func TestQueryParameter(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8000/welcome?name=JOHN", nil)
	record := httptest.NewRecorder()
	lrnwebgo.WelcomeProfile(record, req)

	resp := record.Result()
	body, _ := io.ReadAll(resp.Body);

	assert.Equal(t, "WELCOME JOHN", string(body), "TEST")
	assert.Equal(t, 200, resp.StatusCode, "TEST")
}

func TestMultipleQueryParameter(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8000/welcome?name=FRANKY&name=LUFFY", nil)
	record := httptest.NewRecorder()
	lrnwebgo.WelcomeProfile(record, req)

	resp := record.Result()
	body, _ := io.ReadAll(resp.Body);

	assert.Equal(t, "WELCOME FRANKY LUFFY", string(body), "TEST")
	assert.Equal(t, 200, resp.StatusCode, "TEST")
}





