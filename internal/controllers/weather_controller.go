package controllers

import (
	"fmt"
	"net/http"
)

type WeatherforecastController struct {
	Controller
}

func NewWeatherforecaseController() WeatherforecastController {
	controller := WeatherforecastController{Controller: Controller{Name: "weatherforecast"}}
	controller.MapEndpoint("/cities", controller.GetCities)
	controller.MapEndpoint("/weathers", controller.GetWeathers)

	return controller
}

func (c WeatherforecastController) GetHandlerFuncs() map[string]http.HandlerFunc {
	handlers := make(map[string]http.HandlerFunc)
	controllerPrefix := "/weatherforecast"
	handlers[controllerPrefix+"/cities"] = c.GetCities
	handlers[controllerPrefix+"/weathers"] = c.GetWeathers

	return handlers
}

func (c *WeatherforecastController) GetCities(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Headers", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "*")
	// w.Header().Set("Access-Control-Allow-Credentials", "true")
	fmt.Fprintln(w, "[\"Nanjing\", \"Shanghai\"]")
}

func (c *WeatherforecastController) GetWeathers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "[\"Cloudy\", \"Shiny\"]")
}
