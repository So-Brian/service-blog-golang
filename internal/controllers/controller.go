package controllers

import (
	"fmt"
	"net/http"
)

type IController interface {
	GetHandlerFuncs() map[string]http.HandlerFunc
}

type Controller struct {
	Name     string
	Handlers map[string]http.HandlerFunc
}

func (c *Controller) MapEndpoint(route string, handler http.HandlerFunc) {
	if c.Handlers == nil {
		c.Handlers = make(map[string]http.HandlerFunc)
	}

	endpoint := fmt.Sprintf("/%s%s", c.Name, route)
	c.Handlers[endpoint] = handler
}

func (c Controller) GetHandlerFuncs() map[string]http.HandlerFunc {
	return c.Handlers
}
