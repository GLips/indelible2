package controllers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"

	"github.com/GLips/Indelible2/app/models"

	"github.com/revel/revel"
)

func (c *Controller) GetJSONParam(name string, value interface{}) bool {
	err := json.Unmarshal([]byte(c.JSONParams[name]), &value)
	model, ok := value.(models.Model)
	if ok {
		model.PostProcess()
	} else {
		revel.WARN.Printf("Could not make assertion that %T is a model.", value)
	}
	return err == nil
}

func (c *Controller) PreprocessJSON() revel.Result {
	c.JSONParams = make(map[string]string)
	t := c.Request.Header.Get("Content-Type")
	if strings.Contains(t, "text/json") ||
		strings.Contains(t, "application/json") {

		maxJSONData := int64(10 << 20)
		reader := io.LimitReader(c.Request.Body, maxJSONData+1)
		b, err := ioutil.ReadAll(reader)
		if err != nil {
			return nil
		}

		if int64(len(b)) > maxJSONData {
			return nil
		}
		c.Body = b

		var params map[string]*json.RawMessage
		if err := json.Unmarshal([]byte(b), &params); err != nil {
			// If there was an error parsing the JSON body, then just leave the
			// JSONParams as the empty map.
			return nil
		}

		for k, v := range params {
			if v != nil {
				c.JSONParams[k] = string(*v)
			}
		}
	}
	return nil
}

// RenderJSON calls the revel supplied RenderJson. It is supplied to allow for a
// more consistent naming convention of JSON related methods.
func (c *Controller) RenderJSON(i interface{}) revel.Result {
	return c.RenderJson(i)
}

func (c *Controller) RenderJSONOk() revel.Result {
	return c.RenderJSON("ok")
}

func (c *Controller) RenderJSONValidation() revel.Result {
	if !c.Validation.HasErrors() {
		return c.RenderJSONOk()
	}

	errors := make(map[string][]string)
	for _, message := range c.Validation.ErrorMap() {
		errors[message.Key] = []string{message.Message}
	}

	c.Response.Status = 422
	return c.RenderJSON(map[string]map[string][]string{"errors": errors})
}
