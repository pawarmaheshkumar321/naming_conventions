package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"golang.org/x/net/http2"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		logRequestDetails(r)
		fmt.Fprintf(w, "Handling Incoming Orders")
	})

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		logRequestDetails(r)
		fmt.Fprintf(w, "Handling Users")
	})

	port := 3000

	//load cert and key
	//openssl req -x509 -newkey rsa:2048 -nodes -keyout key.pem -out cert.pem -days 365
	//openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem -config openssl.cnf
	cert := "cert.pem"
	key := "key.pem"

	//configure TLS
	tlsConfig := &tls.Config{
		MinVersion: tls.VersionTLS12,
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  loadClientCAs(),
	}

	//create a custom server
	server := &http.Server{
		Addr:      fmt.Sprintf(":%d", port),
		Handler:   nil,
		TLSConfig: tlsConfig,
	}

	err := http2.ConfigureServer(server, &http2.Server{})
	if err != nil {
		log.Fatalln("Coluld not start http2 server", err)
	}

	fmt.Println("Server is running on port:", port)

	err = server.ListenAndServeTLS(cert, key)
	if err != nil {
		log.Fatalln("Coluld not start server", err)
	}

}

func logRequestDetails(r *http.Request) {
	httpVersion := r.Proto
	fmt.Println("Received reuest with HTTP version", httpVersion)

	if r.TLS != nil {
		tlsVersion := getTLSVersionName(r.TLS.Version)
		fmt.Println("Received request with TLS version", tlsVersion)
	} else {
		fmt.Println("Received request without TLS")
	}
}

func getTLSVersionName(version uint16) string {

	switch version {
	case tls.VersionTLS10:
		return "TLS 1.0"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS12:
		return "TLS 1.2"

	default:
		return "Unknown TLS version found"

	}
}

func loadClientCAs() *x509.CertPool {
	clientCAs := x509.NewCertPool()
	caCert, err := os.ReadFile("cert.pem")
	if err != nil {
		log.Fatalln("Coluld not load client CA:", err)
	}
	clientCAs.AppendCertsFromPEM(caCert)
	return clientCAs
}
