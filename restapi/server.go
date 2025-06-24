package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello Server!!!")
	})

	const serverAddr string = "127.0.0.1:3000"
	//const port string = ":3000"

	fmt.Println("Server is listening on port 3000")
	err := http.ListenAndServe(serverAddr, nil)
	if err != nil {
		log.Fatalln("Error starting server", err)
	}

}
