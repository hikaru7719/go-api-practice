package main

import (
	"net/http"
	"fmt"
	"html"
	"log"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	fmt.Fprintf(w ,"hello ,%q", html.EscapeString(r.URL.Path))
}

func main() {
	router := httprouter.New()
	router.GET("/:path",Index)
	log.Fatal(http.ListenAndServe(":8080",router))
}
