package main

import (
	"github.com/acmacalister/goatee"
	"html/template"
	"log"
	"net/http"
)

func main() {
	go httpServer()

	err := goatee.CreateServer([]string{"chan1", "chan2"})

	if err != nil {
		log.Fatal("Error: ", err.Error())
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}

func httpServer() {
	log.Print("starting http server...")
	http.HandleFunc("/test", testHandler)
	log.Fatal(http.ListenAndServe(":1236", nil))
}
