package main

import (
	"net/http"
	"log"
	"github.com/julienschmidt/httprouter"
)





func main() {
	router := httprouter.New()
	router.GET("/",Logging(Index,"index"))
	router.GET("/todos",Logging(TodoIndex,"todo-index"))
	router.GET("/todos/:todoId", Logging(TodoShow,"todo-show"))

	log.Fatal(http.ListenAndServe(":8080",router))
}
