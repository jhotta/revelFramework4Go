package controllers

import (
	"fmt"
	"github.com/robfig/revel"
	"revelFramework4Go/sampleBlogSite/app/models"

)

type Posts struct {
	App
}

func (c Posts) Index() revel.Result {
	results, err := c.Txn.Get(models.Post{},2)
	if err != nil {
		panic(err)
	}
	payLoad := results.(*models.Post).Body
	return c.Render(payLoad)
}

func (c Posts) Index2() revel.Result {
	results, err := c.Txn.Select(models.Post{},
		`select * from Post limit ?, ?`, 0, 10)
	if err != nil {
		panic(err)
	}

	var p []*models.Post
	for _, r := range results {
		p = append(p, r.(*models.Post))
	}

	payLoad := p
	fmt.Println(payLoad)
	return c.Render(payLoad)
}