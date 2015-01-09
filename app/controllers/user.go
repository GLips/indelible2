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
