package lrnwebgo

import (
	"net/http"
	"lrnwebgo"
)


func main()  {
	http.HandleFunc("/", lrnwebgo.Home);

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}