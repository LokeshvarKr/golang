package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	r.HandleFunc("/hello-world", helloWorld)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":5000", r))
}

func handler(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Create("data/test_data1.txt")
	f.WriteString("Hello from Go")
	f.Sync()

	fmt.Fprint(w, "Hello from Go")
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	f, _ := os.Create("data/test_data2.txt")
	f.WriteString("Hello World.............")
	f.Sync()

	fmt.Fprint(w, "Hello World.............")
}
