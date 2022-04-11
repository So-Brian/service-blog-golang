package main

import (
	"net/http"

	"github.com/so-brian/service-blog-golang/internal/host"
)

func main() {
	host := &host.Host{}
	http.ListenAndServe("127.0.0.1:8080", host)
}
