package entry

import (
	"github.com/GLips/Indelible2/app/db"
	"github.com/GLips/Indelible2/app/models"
)

type Entry struct {
	models.Tmpl
	Content        string `sql:"type:text;"`
	UserId         int64
	SecondsToWrite int
}

func (e Entry) One() string {
	return "entry"
}

func (e Entry) Many() string {
	return "entries"
}

func (e Entry) CheckCreate() {
}

func (e *Entry) Create() {
	connection := db.New()
	connection.Save(e)
}

func (e Entry) Delete() {
}

func (e Entry) Update() {
}
