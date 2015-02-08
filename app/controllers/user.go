package controllers

import (
	"github.com/GLips/Indelible2/app/models/user"
	"github.com/revel/revel"
)

// userSessionKey is the session key that holds the user name current logged in
// user. If there is no value associated with this key, then the current session
// is not logged in.
const userSessionKey = "username"

type User struct {
	Controller
}

func (c User) Create() revel.Result {
	u := &user.User{}
	c.GetJSONParam(u.One(), u)

	u.Validate(c.Validation)

	return c.basicCreate(u)
}

func (c User) Login() revel.Result {
	var auth user.User
	c.GetJSONParam("user", &auth)

	user, err := auth.Login(c.Validation)
	if err {
		return c.RenderJSONValidation()
	} else {
		// Associate the current session with the given user name.
		c.Session[userSessionKey] = user.Username

		// TODO: Provide a user JSON for logged in users to
		//       bootstrap Ember on initial page load.
		user.Strip()
		return c.RenderJSON(map[string]interface{}{user.One(): user})
	}
}

// CheckedLoggedIn is an interceptor that will determine if the current request
// is from a logged in user or not. It will populate the LogeedIn and Username
// fields of the controller appropriately.
func (c *Controller) CheckLoggedIn() revel.Result {
	c.CurrentUsername = c.Session[userSessionKey]
	c.LoggedIn = (c.CurrentUsername != "")
	return nil
}
