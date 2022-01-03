package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", "9999"),
		Handler: setupRouter(),
	}
	server.ListenAndServe()
}

func setupRouter() *mux.Router {
	router := mux.NewRouter()
	router.Methods(http.MethodPost).Path("/echo").HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println(r.RemoteAddr)
		rw.WriteHeader(200)
		rw.Write([]byte("echo\r\n"))
	})

	return router
}
