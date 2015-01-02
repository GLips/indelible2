package controllers

import (
	"github.com/GLips/Indelible2/app/db"

	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

type Controller struct {
	*revel.Controller
	DBSess     gorm.DB
	JSONParams map[string]string
	Body       []byte
}

func init() {
	revel.InterceptMethod((*Controller).DBBegin, revel.BEFORE)
	revel.InterceptMethod((*Controller).PreprocessJSON, revel.BEFORE)
}

func (c *Controller) DBBegin() revel.Result {
	c.DBSess = db.New()
	return nil
}
