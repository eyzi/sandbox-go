package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func startRouter() http.Handler {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homepage).Methods("GET")
	return myRouter
}

func startTcp(port string, handler http.Handler) {
	host := "0.0.0.0"
	if port == "" {
		port = "80"
	}
	addr := host + ":" + port
	println("server is running on http://" + addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}

/*
	creates a TCP server using env variable `PORT`
	as server port or defaults to port 80
*/
func CreateServer() {
	handler := startRouter()
	startTcp(os.Getenv("PORT"), handler)
}
