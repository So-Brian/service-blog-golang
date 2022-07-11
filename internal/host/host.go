package host

import (
	"fmt"
	"net/http"
	"regexp"
)

type Host struct {
	handlerFuncs map[string]http.HandlerFunc
}

func NewHost(m map[string]http.HandlerFunc) *Host {
	host := &Host{handlerFuncs: m}

	return host
}

func (host *Host) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for k, f := range host.handlerFuncs {
		re := regexp.MustCompile(k)
		match := re.FindStringSubmatch(r.URL.Path)

		if len(match) == 2 {
			f(w, r)

			return
		}

	}

	if f, ok := host.handlerFuncs[r.URL.Path]; ok {
		f(w, r)
	} else {
		fmt.Fprintln(w, "Hello world by host")
	}
}
