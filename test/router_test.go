package test

import ( 
	"testing"
	"net/http/httptest"
	"net/http"
	"io"
	"fmt"
	"lrnwebgo"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T)  {
	

	request := httptest.NewRequest("GET", "http://localhost:8000/", nil)
	recorder := httptest.NewRecorder()

	router := httprouter.New()
	router.GET("/", lrnwebgo.Index)
	router.ServeHTTP(recorder, request);

	response := recorder.Result()


	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "HELLO", string(body), "TEST")
	assert.Equal(t, 200, response.StatusCode, "TEST")
}


func TestDetail(t *testing.T) {
	req, err := http.NewRequest("GET", "/welcome/tester", nil)
	if err != nil {
		panic(err)
	}
	res := httptest.NewRecorder()

	router := httprouter.New()
	router.GET("/welcome/:name", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		lrnwebgo.Detail(w, r, params)
	})

	router.ServeHTTP(res, req);

	response := res.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "WELCOME tester", string(body), "TEST")
	assert.Equal(t, 200, response.StatusCode, "OKEE")
}

func TestCatchAllParameter(t *testing.T)  {
	req, err := http.NewRequest("GET", "/images/small/dummy.jpeg", nil)
	if err != nil {
		panic(err)
	}
	res := httptest.NewRecorder()

	router := httprouter.New()
	router.GET("/images/*image", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Fprint(w, "Image ", params.ByName("image"));	
	})

	router.ServeHTTP(res, req);

	response := res.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Image /small/dummy.jpeg", string(body), "TEST")
	assert.Equal(t, 200, response.StatusCode, "OKEE")	
}

func TestPanic(t *testing.T)  {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		panic(err)
	}
	res := httptest.NewRecorder()

	router := httprouter.New()
	router.PanicHandler = func(w http.ResponseWriter, r *http.Request, error interface{}) {
		fmt.Fprint(w, "Panic ", error)
	}

	router.GET("/", func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		panic("ERROR")
	})

	router.ServeHTTP(res, req);

	response := res.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "Panic ERROR", string(body), "TEST")
}

func TestRouterNotFound(t *testing.T)  {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		panic(err)
	}
	res := httptest.NewRecorder()

	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func( w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "PRINT NOT FOUND")
	});

	router.ServeHTTP(res, req);
	response := res.Result()
	body, _ := io.ReadAll(response.Body)
	assert.Equal(t, "PRINT NOT FOUND", string(body), "TEST")
}