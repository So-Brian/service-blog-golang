package host

import (
	"fmt"
	"net/http"
)

type Host struct {
}

func (host *Host) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world by host")
}
