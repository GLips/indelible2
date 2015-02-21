package controllers

import (
	"github.com/GLips/Indelible2/app/db"
	"github.com/GLips/Indelible2/app/models/entry"
	"github.com/revel/revel"
)

type Entry struct {
	Controller
}

func (c Entry) Create() revel.Result {
	d := db.New()
	d.DropTable(&entry.Entry{})
	d.CreateTable(&entry.Entry{})
	if c.IsLoggedIn() {
		e := &entry.Entry{}
		u := c.ActiveUser()

		e.UserId = u.Id

		c.GetJSONParam(e.One(), e)
		return c.basicCreate(e)
	} else {
		return c.RenderJSONError("Looks like you're not logged in, try logging in before you save this entry.")
	}
}
