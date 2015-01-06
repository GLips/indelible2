package controllers

import (
	"github.com/GLips/Indelible2/app/models"

	"github.com/revel/revel"
)

type strippableModel interface {
	models.Model
	Strip()
}

type basicCreateModel interface {
	models.Model
	CheckCreate()
}

type basicUpdateModel interface {
	models.Model
	CheckUpdate()
}

type basicDeleteModel interface {
	models.Model
	CheckDelete()
}

type basicGetModel interface {
	models.Model
	CheckGet()
}

func (c *Controller) basicCreate(model basicCreateModel) revel.Result {
	model.CheckCreate()
	if c.Validation.HasErrors() {
		return c.RenderJSONValidation()
	}

	model.Create()

	strippable, ok := model.(strippableModel)
	if ok {
		strippable.Strip()
	}
	return c.RenderJSON(map[string]interface{}{model.One(): model})
}
