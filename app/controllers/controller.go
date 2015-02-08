package controllers

import (
	"sync"

	"github.com/GLips/Indelible2/app/db"
	"github.com/GLips/Indelible2/app/models/user"

	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

var once sync.Once

type Controller struct {
	*revel.Controller
	DBSess gorm.DB

	// The logged in state of the request (user.go)
	LoggedIn        bool
	CurrentUsername string
	CurrentUser     user.User

	JSONParams map[string]string
	Body       []byte
}

func (c *Controller) IsLoggedIn() bool {
	return c.LoggedIn
}

func (c *Controller) ActiveUsername() string {
	return c.CurrentUsername
}

func (c *Controller) ActiveUser() user.User {
	if c.IsLoggedIn() {
		// Only hit the DB to try to find the currently logged in user once.
		once.Do(func() {
			c.CurrentUser = user.FindByUsername(c.ActiveUsername())
		})
	}
	return c.CurrentUser
}

func init() {
	revel.InterceptMethod((*Controller).DBBegin, revel.BEFORE)
	revel.InterceptMethod((*Controller).PreprocessJSON, revel.BEFORE)
	// Grab the current user's info from the session so it's accessible
	// to the controller.
	revel.InterceptMethod((*Controller).CheckLoggedIn, revel.BEFORE)

	// Export the user to the template as 'user' for bootstrapping Ember
	revel.InterceptMethod((*Controller).ExportUser, revel.AFTER)
}

func (c *Controller) DBBegin() revel.Result {
	c.DBSess = db.New()
	return nil
}
