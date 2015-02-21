package controllers

import (
	"github.com/GLips/Indelible2/app/models/entry"
	"github.com/revel/revel"
)

type Entry struct {
	Controller
}

func (c Entry) Create() revel.Result {
	if c.IsLoggedIn() {
		e := &entry.Entry{}
		u := c.ActiveUser()

		e.UserId = u.Id

		// TODO: Create a user entry index page
		c.GetJSONParam(e.One(), e)
		return c.basicCreate(e)
	} else {
		return c.RenderJSONError("Looks like you're not logged in, try logging in before you save this entry.")
	}
}
