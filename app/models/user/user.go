package user

import (
	"github.com/GLips/Indelible2/app/db"
	"github.com/GLips/Indelible2/app/models"

	"code.google.com/p/go.crypto/bcrypt"

	"github.com/revel/revel"
)

type User struct {
	models.Tmpl
	Identifier      string `sql:"size:128"`
	Email           string
	Username        string `sql:"size:32"`
	Password        string `sql:"-"`
	PasswordConfirm string `sql:"-"`
	PasswordHash    string `sql:"size:128"`
}

func (u User) One() string {
	return "user"
}

func (u User) Many() string {
	return "users"
}

func (u User) CheckCreate() {
}

func (u *User) Create() {
	connection := db.New()
	connection.Save(u)
}

func (u User) Update() {
}

func (u User) Delete() {
}

func (u *User) Validate(v *revel.Validation) {
	u.PasswordHash = ""

	// Checks on the username
	v.Required(u.Username).
		Key("Username").
		Message("We need a username to create your account.")
	v.MinSize(u.Username, 3).
		Key("Username").
		Message("A username should be at least 3 characters long.")
	v.MaxSize(u.Username, 32).
		Key("Username").
		Message("A username should be less than 32 characters long.")

	u.validatePassword(v)
}

func (u *User) validatePassword(v *revel.Validation) {
	v.Required(u.Password).
		Key("Password").
		Message("A password must be supplied.")
	v.MinSize(u.Password, 8).
		Key("Password").
		Message("Passwords must be at least 8 characters.")
	//v.Check(u.Password, models.EqualTo{u.PasswordConfirm}).
	//	Key("PasswordConfirm").
	//	Message("Passwords do not match.")

	// If there are no errors, generate a new password hash.
	if !v.HasErrors() {
		p := []byte(u.Password)
		pwHash, _ := bcrypt.GenerateFromPassword(p, bcrypt.DefaultCost)
		u.PasswordHash = string(pwHash)
		//v.Check(err, models.EqualTo{nil}).
		//	Key(app.FlashErrorKey).
		//	Message("An error occurred, please try again.")
	}

	u.Password = ""
	u.PasswordConfirm = ""
}
