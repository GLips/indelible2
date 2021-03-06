package controllers

import (
	"github.com/GLips/Indelible2/app/db"
	"github.com/GLips/Indelible2/app/models/entry"
	"github.com/revel/revel"
)

type Entry struct {
	Controller
}

func (c Entry) Index() revel.Result {
	if c.IsLoggedIn() {
		// TODO: Abstract this out into a basicQuery function
		// TODO: Hook up an ESP
		// TODO: Send emails to welcome users after registering
		// We initialize entries like this so RenderJSON renders a
		// blank array as [] instead of null.
		entries := make([]entry.Entry, 0)
		var e entry.Entry
		currentUser := c.ActiveUser()
		connection := db.New()
		connection.Model(&currentUser).Order("created_at desc").Limit(10).Related(&entries)
		return c.RenderJSON(map[string]interface{}{e.Many(): entries})
	} else {
		return c.RenderJSONError("You need to login first to retrieve a list of your entries.")
	}
}

func (c Entry) Create() revel.Result {
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
