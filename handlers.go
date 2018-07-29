package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
	"html"
	"encoding/json"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	fmt.Fprintf(w ,"hello ,%q", html.EscapeString(r.URL.Path))
}

func TodoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	todos := Todos{
		Todo{Name: "Write presentation"},
		Todo{Name: "Host meetup"},
	}
	json.NewEncoder(w).Encode(todos)
}

func TodoShow(w http.ResponseWriter, t *http.Request, ps httprouter.Params){
	fmt.Fprintf(w, "Todo show: %s", ps.ByName("todoId"))
}