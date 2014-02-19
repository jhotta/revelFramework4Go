package controllers

import (
	"github.com/robfig/revel"
	"revelFramework4Go/sampleBlogSite/app/models"
)

type App struct {
	GorpController
}

func (c App) Index() revel.Result {
	title := "HOME"
	return c.Render(title)
}

func (c App) Hello(myName string) revel.Result {
	c.Validation.Required(myName).Message("Your name is required!")
	c.Validation.MinSize(myName, 3).Message("Your name is not long enough!")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	post := &models.Post{0, myName}
	if err := Dbm.Insert(post); err != nil {
		panic(err)
	}

	title := "HELLO"
	return c.Render(title, myName)
}

func (c App) List() revel.Result {
	title := "LIST"
	return c.Render(title)
}

func (c App) Project() revel.Result {
	title := "PROJECT"
	return c.Render(title)
}