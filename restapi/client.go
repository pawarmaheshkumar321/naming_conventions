package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	client := &http.Client{}

	resp, err := client.Get("http://tilltoss.net:5000/tahsil/getsettings2")
	if err != nil {
		fmt.Println("Error making GET request", err)
		return
	}

	defer resp.Body.Close()

	// read and print response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading Response body", err)
		return
	}

	fmt.Printf("Response received from server is %v", string(body))
}
