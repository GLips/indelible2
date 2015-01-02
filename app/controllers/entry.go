package controllers

import (
	"github.com/GLips/Indelible2/app/models/entry"
	"github.com/revel/revel"
)

type Entry struct {
	Controller
}

func (c Entry) Create() revel.Result {
	val := entry.Entry{}
	c.GetJSONParam(val.One(), &val)
	val.Save()
	return c.RenderJson(map[string]interface{}{val.One(): val})
}
