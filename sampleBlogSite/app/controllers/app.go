package controllers

import "github.com/robfig/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {

	payLoad := map[string] string {
		"title": "HOME",
	}
	return c.Render(payLoad)
}


func (c App) Hello(myName string) revel.Result {
	c.Validation.Required(myName).Message("Your name is required!")
	c.Validation.MinSize(myName, 3).Message("Your name is not long enough!")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	payLoad := map[string] string {
		"title": "HELLO",
		"myName": myName,
	}

	return c.Render(payLoad)
}

func (c App) List() revel.Result {

	payLoad := map[string] string {
		"title": "LIST",
	}
	return c.Render(payLoad)
}

func (c App) Project() revel.Result {

	payLoad := map[string] string {
		"title": "PROJECT",
	}
	return c.Render(payLoad)
}