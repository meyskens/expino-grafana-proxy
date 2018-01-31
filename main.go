package main

import (
	"encoding/base64"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var credentials string

func main() {
	credentials = base64.StdEncoding.EncodeToString([]byte(os.Getenv("USERNAME") + ":" + os.Getenv("PASSWORD")))
	remote, err := url.Parse(os.Getenv("GRAFANAURL"))
	if err != nil {
		panic(err)
	}

	proxy := httputil.NewSingleHostReverseProxy(remote)
	http.HandleFunc("/", handler(proxy))
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}

func handler(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		r.Header.Set("Authorization", "Basic "+credentials)
		p.ServeHTTP(w, r)
	}
}
