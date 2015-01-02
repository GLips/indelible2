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
