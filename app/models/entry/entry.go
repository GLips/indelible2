package entry

import (
	"github.com/GLips/Indelible2/app/db"
	"github.com/GLips/Indelible2/app/models"
)

type Entry struct {
	models.Tmpl
	Content string
}

func (e Entry) One() string {
	return "entry"
}

func (e Entry) Many() string {
	return "entries"
}

func (e *Entry) Save() {
	connection := db.New()
	e.Touch()
	connection.Save(e)
}

func (e Entry) Delete() {
}
