package models

import (
	"github.com/robfig/revel"
)

type Post struct {
	PostId int
	Body   string
}

func (post *Post) Validate(v *revel.Validation) {

	v.MaxSize(post.Body, 100)

}
