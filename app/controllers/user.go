package controllers

import (
	"github.com/GLips/Indelible2/app/models/user"
	"github.com/revel/revel"
)

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
		// TODO: Add a persistent feature to track logged in users.
		// TODO: Provide a user JSON for logged in users to
		//       bootstrap Ember on initial page load.
		user.Strip()
		return c.RenderJSON(map[string]interface{}{user.One(): user})
	}
}
