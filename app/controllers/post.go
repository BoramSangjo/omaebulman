package controllers

import (
	"github.com/revel/revel"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/BoramSangjo/omaebulman/app/models"
	"github.com/BoramSangjo/omaebulman/app/routes"
	"net/url"
	"log"
)

type Post struct {
	App
}

func (c Post) Index() revel.Result {
	var posts []models.Post
	res, err := http.Get("http://60.0.1.47:3000/board")
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(b, &posts)
	return c.Render(posts)
}

func (c Post) New() revel.Result {
	return c.Render()
}

func (c Post) Create(title, description string) revel.Result {
	res, err := http.PostForm("http://60.0.1.47:3000/board", url.Values{"title":{title}, "description":{description}})
	if err != nil {
		panic(err)
	}
	if res.StatusCode == 500 {
		log.Fatal(http.StatusText(500))
	}
	return c.Redirect(routes.Post.Index())
}
