package controllers

import (
	"github.com/revel/revel"
	"github.com/BoramSangjo/omaebulman/app/routes"
	"net/http"
	"net/url"
	"log"
)

type Login struct {
	App
}

func (c Login) Index() revel.Result {
	return c.Render()
}

func (c Login) New(id, password string) revel.Result {
	_, err := http.PostForm("http://60.0.1.47:3000/auth/login", url.Values{"name":{id}, "pw":{password}})
	if err != nil {
		log.Fatal(err)
	}
	return c.Redirect(routes.App.Index())
}
