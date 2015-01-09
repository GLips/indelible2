package controllers

import (
	"github.com/GLips/Indelible2/app/models/entry"
	"github.com/revel/revel"
)

type Entry struct {
	Controller
}

func (c Entry) Create() revel.Result {
	e := &entry.Entry{}
	c.GetJSONParam(e.One(), e)
	return c.basicCreate(e)
}
