package lrnwebgo

import (
	"fmt"
	"net/http"
	"strings"
	"github.com/julienschmidt/httprouter"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("EXECUTED")
	fmt.Fprint(w, "HOME PAGE")
	fmt.Println("AF EXECUTED")
} 


func Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "WELCOME")
}

func WelcomeProfile(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	if strings.Join(params["name"], "") == "" {
		fmt.Fprint(w, "WELCOME ")
		return
	}
	fmt.Fprint(w, "WELCOME ", strings.Join(params["name"], " "))
}

func GetRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, r.Method)
}


func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	fmt.Fprint(w, "HELLO");	
}

func Detail(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	fmt.Fprint(w, "WELCOME ", param.ByName("name"));	
}