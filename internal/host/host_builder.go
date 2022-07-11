package host

import (
	"net/http"

	"github.com/so-brian/service-blog-golang/internal/controllers"
)

type HostBuilder struct {
	handlerFuncs map[string]http.HandlerFunc
}

func WebHostBuilder() *HostBuilder {
	builder := &HostBuilder{}
	builder.handlerFuncs = make(map[string]http.HandlerFunc)

	return builder
}

func (builder *HostBuilder) Configure(f func(builder *HostBuilder)) *HostBuilder {
	f(builder)

	return builder
}

func (builder *HostBuilder) Build() *Host {
	host := NewHost(builder.handlerFuncs)

	return host
}

func (builder *HostBuilder) AddControllers(controllers ...controllers.IController) *HostBuilder {
	for _, c := range controllers {
		for k, v := range c.GetHandlerFuncs() {
			builder.handlerFuncs[k] = v
		}
	}

	return builder
}
