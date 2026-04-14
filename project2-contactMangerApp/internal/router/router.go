package router

import "net/http"

/*Connects routes → handlers*/
func SetupRoutes() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/hello", helloHandler)
}
