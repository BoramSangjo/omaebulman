package controllers

import (
	"github.com/revel/revel"
	"github.com/BoramSangjo/omaebulman/app/models"
)

type App struct {
	*revel.Controller
	CurrentUser *models.User
}

func (c App) Index() revel.Result {
	return c.Render()
}
