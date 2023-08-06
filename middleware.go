package lrnwebgo

import (
	"net/http"
	"fmt"
)
type LogMiddleware struct {
	Handler http.Handler
}

type HandleError struct {
	Handler http.Handler
}

func (handleErr *HandleError) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	defer func ()  {
		err := recover()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "ERROR ", err)
		}
	}()
	handleErr.Handler.ServeHTTP(w, r)
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	middleware.Handler.ServeHTTP(w, r)
}


func Action()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", Home)

	logMiddleware := &LogMiddleware {
		Handler : mux,
	}

	errorHandler := &HandleError {
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

