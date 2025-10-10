package main

import (
	"fmt"
	"net/http"
)

// APIS
func handleHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Server apied created....")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handleHello)

	//Server Config
	fmt.Println("Server is running on port, 3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("ERROR server running ", err.Error())
	}
}
