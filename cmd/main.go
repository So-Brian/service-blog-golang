package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/so-brian/service-blog-golang/internal/controllers"
	"github.com/so-brian/service-blog-golang/internal/host"
)

func main() {
	port := os.Getenv("port")
	endpoint := fmt.Sprintf(":%s", port) //os.Args[1]

	host := host.WebHostBuilder().Configure(configure).Build()
	fmt.Printf("Start listening on %s", endpoint)
	http.ListenAndServe(endpoint, host)
}

func configure(builder *host.HostBuilder) {
	builder.AddControllers(
		controllers.NewWeatherforecaseController(),
		controllers.NewInternalController(),
		controllers.NewBlogController())
}
