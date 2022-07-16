package controllers

import (
	"github.com/labstack/echo/v4"
	"go-checkin/utils/session"
)

type HomeController struct {
	BaseBackendController
}

func NewHomeController() HomeController {
	return HomeController{
		BaseBackendController: BaseBackendController{
			Menu:        "Home",
			BreadCrumbs: []map[string]interface{}{},
		},
	}
}

func (c *HomeController) Index(ctx echo.Context) error {
	breadCrumbs := map[string]interface{}{
		"menu": "Home",
		"link": "/check/home",
	}
	userInfo, _ := session.Manager.Get(ctx, session.SessionId)
	return Render(ctx, "Home", "index", c.Menu, session.GetFlashMessage(ctx),
		append(c.BreadCrumbs, breadCrumbs), userInfo)
}
