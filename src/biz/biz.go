package main

import (
	"net/http"
)

func helloHandler(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Scada.Data.Server"))
}

func getHandler(response http.ResponseWriter, request *http.Request) {

	response.Write([]byte("Scada.Data.Server"))
}

func setHandler(response http.ResponseWriter, request *http.Request) {

	response.Write([]byte("Scada.Data.Server"))
}

func queryHandler(response http.ResponseWriter, request *http.Request) {

	response.Write([]byte("Scada.Data.Server"))
}

func main() {
	//http.Handle("/css/", http.FileServer(http.Dir("template")))
	//http.Handle("/js/", http.FileServer(http.Dir("template")))

	http.HandleFunc("/hello", helloHandler)

	http.HandleFunc("/get", getHandler)
	http.HandleFunc("/set", setHandler)

	http.HandleFunc("/query", queryHandler)

	http.ListenAndServe(":8888", nil)

}
