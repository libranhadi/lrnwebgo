package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"lrnwebgo"
)


func main()  {
	router := httprouter.New()
	router.GET("/", lrnwebgo.Index)
	router.GET("/welcome/:name", lrnwebgo.Detail)
	log.Fatal(http.ListenAndServe(":8000", router))
}

