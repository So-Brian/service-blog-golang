package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/so-brian/service-blog-golang/internal/repositories"
)

type BlogController struct {
	Controller
}

func NewBlogController() BlogController {
	controller := BlogController{Controller: Controller{Name: "blog"}}
	controller.MapEndpoint("/", controller.GetBlogs)

	return controller
}

func (c BlogController) GetHandlerFuncs() map[string]http.HandlerFunc {
	handlers := make(map[string]http.HandlerFunc)
	controllerPrefix := "/blog"
	handlers[controllerPrefix+"/"] = c.GetBlogs
	handlers[controllerPrefix+"/(?P<id>\\d+)"] = c.GetBlog

	return handlers
}

func (c BlogController) GetBlogs(w http.ResponseWriter, r *http.Request) {
	repo, _ := repositories.NewBlogRepository()
	blogs, _ := repo.GetBlogs()

	res, err := json.Marshal(blogs)

	if err != nil {
		fmt.Println(err)
	}

	// json := fmt.Sprintln(string(res))
	// fmt.Fprintln(w, json)

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func (c BlogController) GetBlog(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile(`/blog/(?P<id>\d+)`)
	match := re.FindStringSubmatch(r.URL.Path)
	id, err := strconv.Atoi((match[1]))

	if err != nil {
		log.Fatal(err.Error())
	}

	repo, err := repositories.NewBlogRepository()
	if err != nil {
		fmt.Fprintln(w, err.Error())

		return
	}

	blog, err := repo.GetBlog(id)
	if err != nil {
		fmt.Fprintln(w, err.Error())

		return
	}

	res, err := json.Marshal(blog)

	if err != nil {
		fmt.Fprintln(w, err.Error())

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
