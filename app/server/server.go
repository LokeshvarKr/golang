package main

import "net/http"

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome !"))
}
func student(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(200)
	w.Write([]byte("Student Section"))
}
func teacher(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(200)
	w.Write([]byte("Teacher Section"))
}
func main() {
	http.HandleFunc("/", welcome)
	http.HandleFunc("/student", student)
	http.HandleFunc("/teacher", teacher)
	http.ListenAndServe("localhost:8080", nil)
}
