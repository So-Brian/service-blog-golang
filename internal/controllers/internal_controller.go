package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

type InternalController struct {
	Controller
}

func NewInternalController() InternalController {
	controller := InternalController{Controller: Controller{Name: "internal"}}
	controller.MapEndpoint("/env", controller.GetEnv)
	controller.MapEndpoint("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hello, world!")
	})

	return controller
}

func (c InternalController) GetEnv(w http.ResponseWriter, r *http.Request) {
	env := ""

	// fetcha all env variables
	for _, element := range os.Environ() {
		variable := strings.Split(element, "=")
		env += fmt.Sprintf("%s =>  %s \n", variable[0], variable[1])
	}

	fmt.Fprintln(w, env)
}
