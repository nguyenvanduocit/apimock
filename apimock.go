package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"flag"
	"fmt"
	"log"
	"html/template"
)

func APIHandler(w http.ResponseWriter, r *http.Request){

	expectedResponse := r.FormValue("expected_response")
	expectedMime := r.FormValue("expected_response_mime")

	if(expectedMime != ""){
		w.Header().Add("Content-Type", expectedMime)
	}

	w.Write([]byte(expectedResponse))
}

func HelpHandler(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("template/help.html")
	t.Execute(w, nil)
}

func main() {
	var ip, port string
	flag.StringVar(&ip, "ip", "", "ip")
	flag.StringVar(&port, "port", "8080", "Port")
	flag.Parse()

	address := fmt.Sprintf("%s:%s", ip, port)
	router := mux.NewRouter()

	router.HandleFunc("/", APIHandler)
	router.HandleFunc("/help", HelpHandler)
	log.Fatal(http.ListenAndServe(address, router))
}
