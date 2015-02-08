// Controller that allows for the storing of user's logged in status.
package controllers

import (
	"encoding/json"
	"strings"

	"github.com/revel/revel"
)

const (
	// userBootstrapTokenParam is parameter that is reserved for embedding a
	// JavaScript representation of the current logged in user in all page loads.
	userBootstrapTokenParam = "user"

	// sessionCookie is the cookie entry that will be used to store the current
	// session id.
	sessionCookie = "session"
)

// ExportUser will make the stripped user struct available to the rendered
// template for bootstrapping in EmberJS. Does a lot of checking to ensure
// it doesn't do unnecessary work.
func (c *Controller) ExportUser() revel.Result {
	const undefined = "{}"
	if !c.IsLoggedIn() {
		c.RenderArgs[userBootstrapTokenParam] = undefined
		return nil
	}

	// We only need to export the user variable on GET requests—even then not on
	// all GET requests, only on ones that are hitting non-/api URLs.
	url := c.Request.URL
	if strings.HasPrefix(url.String(), "/api") {
		revel.INFO.Printf("URL requested is an API endpoint—no need to bootstrap Ember.")
		return nil
	}

	u := c.ActiveUser()
	if u.Id == 0 {
		revel.WARN.Printf("Couldn't find the user account associated with the provided username, '%+v'. Got %+v", c.ActiveUsername(), u)
		c.RenderArgs[userBootstrapTokenParam] = undefined
		return nil
	}

	u.Strip()
	bytes, err := json.Marshal(u)
	if err != nil {
		revel.ERROR.Printf("Couldn't JSONify the user object: %+v", u)
		c.RenderArgs[userBootstrapTokenParam] = undefined
		return nil
	}

	c.RenderArgs[userBootstrapTokenParam] = string(bytes)

	return nil
}
