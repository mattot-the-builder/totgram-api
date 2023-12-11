package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "gorm.io/driver/mysql"
)

func main() {
	r := mux.NewRouter()
	http.Handle("/", r)
	fmt.Println("Starting server at localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
