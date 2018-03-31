package controllers

import (
	"github.com/revel/revel"
	"github.com/BoramSangjo/omaebulman/app/routes"
	"net/http"
	"net/url"
	"log"
	"github.com/BoramSangjo/omaebulman/app/models"
)

type Register struct {
	App
}

func (c Register) Index() revel.Result {
	return c.Render()
}

func (c Register) New(id, password string) revel.Result {
	res, err := http.PostForm("http://60.0.1.47:3000/auth/register", url.Values{"name":{id}, "pw":{password}})
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == 200 {
		u := &models.User{Id: id, Password: password}
		c.CurrentUser = u
	}
	return c.Redirect(routes.Login.Index())
}