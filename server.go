package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Write "Hello, world!" to the response body
	io.WriteString(w, "Hello, world!\n")
}

func runServer(mtls bool) {
	// Set up a /hello resource handler
	http.HandleFunc("/hello", helloHandler)

	// Create a CA certificate pool and add cert.pem to it
	caCert, err := ioutil.ReadFile("cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	if mtls {
		// with mtls
		fmt.Printf("Running with mtls=on\n")

		// Create the TLS Config with the CA pool and enable Client certificate validation
		tlsConfig := &tls.Config{
			ClientCAs:  caCertPool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		}
		tlsConfig.BuildNameToCertificate()

		// Create a Server instance to listen on port 8443 with the TLS config
		server := &http.Server{
			Addr:      ":8443",
			TLSConfig: tlsConfig,
		}

		// Listen to HTTPS connections with the server certificate and wait
		log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
	} else {
		// without mtls
		fmt.Printf("Running with mtls=off\n")

		// Create the TLS Config with the CA pool and disable Client certificate validation
		tlsConfig := &tls.Config{
			ClientCAs: caCertPool,
		}
		tlsConfig.BuildNameToCertificate()

		// Create a Server instance to listen on port 8443 with the TLS config
		server := &http.Server{
			Addr:      ":8443",
			TLSConfig: tlsConfig,
		}

		// Listen to HTTPS connections with the server certificate and wait
		log.Fatal(server.ListenAndServeTLS("cert.pem", "key.pem"))
	}
}

func main() {
	boolPtr := flag.Bool("mtls", false, "should the server expect mtls connections?")
	flag.Parse()

	runServer(*boolPtr)
}
