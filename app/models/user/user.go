package user

import (
	"fmt"

	"github.com/GLips/Indelible2/app/db"
	"github.com/GLips/Indelible2/app/models"

	"code.google.com/p/go.crypto/bcrypt"

	"github.com/revel/revel"
)

type User struct {
	models.Tmpl
	Email           string
	Username        string `sql:"size:32"`
	Password        string `sql:"-" json:",omitempty"`
	PasswordConfirm string `sql:"-" json:",omitempty"`
	PasswordHash    string `sql:"size:128" json:",omitempty"`
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

func (u *User) Strip() {
	u.Password = ""
	u.PasswordHash = ""
	u.PasswordConfirm = ""
}

func FindByField(fieldName string, value string) User {
	var u User
	connection := db.New()
	// It appears our ORM doesn't support double variable concatenation,
	// so we have to build the request up in two methods.
	query := fmt.Sprintf("UPPER(%v) = UPPER(?)", fieldName)
	connection.Where(query, value).First(&u)
	return u
}
func FindByUsername(username string) User {
	return FindByField("username", username)
}
func FindByEmail(email string) User {
	return FindByField("email", email)
}

func (u *User) Validate(v *revel.Validation) {
	u.PasswordHash = ""

	// Checks on the username
	v.Required(u.Username).
		Key("Username").
		Message("We need a username to create your account.")
	v.MaxSize(u.Username, 32).
		Key("Username").
		Message("A username should be less than 32 characters long.")

	u.validatePassword(v)
}

func (u *User) validatePassword(v *revel.Validation) {
	v.Required(u.Password).
		Key("Password").
		Message("We need a password to make your account!")
	v.MinSize(u.Password, 6).
		Key("Password").
		Message("Passwords should be at least 6 characters.")
	v.Check(u.Password, models.EqualTo{u.PasswordConfirm}).
		Key("PasswordConfirm").
		Message("Passwords do not match.")

	// If there are no errors, generate a new password hash.
	if !v.HasErrors() {
		p := []byte(u.Password)
		pwHash, _ := bcrypt.GenerateFromPassword(p, bcrypt.DefaultCost)
		u.PasswordHash = string(pwHash)
	}
}

func (u User) Login(v *revel.Validation) (User, bool) {
	username := u.Username

	dbUser := FindByUsername(username)
	dbUser.ValidateLogin(v, u)
	if v.HasErrors() {
		// If it didn't work for the user name, try again for the email address.
		v.Clear()

		dbUser = FindByEmail(username)
		dbUser.ValidateLogin(v, u)
		if v.HasErrors() {
			u = User{}
			return u, true
		}
	}

	return dbUser, false
}

func (u User) ValidateLogin(v *revel.Validation, user User) {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(user.Password))
	v.Check(err, models.EqualTo{nil}).
		Key("Username").
		Message("Username or password is invalid.")
}
