package test

import(
	"testing"
	"net/http"
	"lrnwebgo"
)

func TestRunServe(t *testing.T) {
	http.HandleFunc("/", lrnwebgo.Home);

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}