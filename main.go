
package main

import (
	"log"
	"net/http"
	"github.com/BearCloud/proj0/api"
	"github.com/gorilla/mux"
)

func main() {
	// Create a new mux for routing api calls
	router := mux.NewRouter()


	//Register our endpoints
	//See /api/api.go
	err := api.RegisterRoutes(router)
	if err != nil {

		//If we can't set up our endpoints, we kill our program (by throwing a global error)
		panic(err.Error())
	}

	//Print log to output, very similar to fmt.Println
	//What are the differences?
	log.Println("starting go server")
	
	http.ListenAndServe(":80", router)

}
