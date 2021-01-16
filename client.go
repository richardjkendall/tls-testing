package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf(",%s,%s", name, elapsed)
}

func sendRequest(mtls bool) {
	// get execution duration
	//defer timeTrack(time.Now(), "sendRequest")

	// Create a CA certificate pool and add cert.pem to it
	caCert, err := ioutil.ReadFile("cert.pem")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	if mtls {
		// using mTLS

		// Read the key pair to create certificate
		cert, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
		if err != nil {
			log.Fatal(err)
		}

		// Create a HTTPS client and supply the created CA pool and certificate
		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs:      caCertPool,
					Certificates: []tls.Certificate{cert},
				},
			},
		}

		// get duration
		defer timeTrack(time.Now(), "sendRequest")

		// Request /hello via the created HTTPS client over port 8443 via GET
		r, err := client.Get("https://localhost:8443/hello")
		if err != nil {
			log.Fatal(err)
		}

		// Read the response body
		defer r.Body.Close()
		_, err2 := ioutil.ReadAll(r.Body)
		if err2 != nil {
			log.Fatal(err)
		}
	} else {
		// not using mTLS

		// Create a HTTPS client and supply the created CA pool and certificate
		client := &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs: caCertPool,
				},
			},
		}

		// get duration
		defer timeTrack(time.Now(), "sendRequest")

		// Request /hello via the created HTTPS client over port 8443 via GET
		r, err := client.Get("https://localhost:8443/hello")
		if err != nil {
			log.Fatal(err)
		}

		// Read the response body
		defer r.Body.Close()
		_, err2 := ioutil.ReadAll(r.Body)
		if err2 != nil {
			log.Fatal(err)
		}
	}

}

func main() {
	boolPtr := flag.Bool("mtls", false, "should the server expect mtls connections?")
	numbPtr := flag.Int("loops", 1, "number of times we should run the loop")
	flag.Parse()

	i := 0
	for i < *numbPtr {
		sendRequest(*boolPtr)
		i++
	}
}
