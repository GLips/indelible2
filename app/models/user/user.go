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

func (u *User) Strip() {
	u.Password = ""
	u.PasswordHash = ""
	u.PasswordConfirm = ""
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
		//v.Check(err, models.EqualTo{nil}).
		//	Key(app.FlashErrorKey).
		//	Message("An error occurred, please try again.")
	}

	u.Password = ""
	u.PasswordConfirm = ""
}

func (u User) Login(v *revel.Validation) (User, bool) {
	badId := false
	connection := db.New()
	connection.Where("UPPER(username) = UPPER(?)", u.Username).First(&u)
	if u.Id == 0 {
		badId = true
	}

	// Check if the given credentials are sufficient for authentication.
	u.ValidateLogin(v, u)
	if v.HasErrors() {
		// If it didn't work for the user name, try again for the email address.
		v.Clear()

		connection.Where("UPPER(email) = UPPER(?)", u.Username).First(&u)
		if u.Id == 0 && badId {
			v.Check(u.Id, models.EqualTo{1}).
				Key("Username").
				Message("Username or email not found. Are you sure you entered your ID correctly?")
			return u, true
		}

		u.ValidateLogin(v, u)
		if v.HasErrors() {
			u = User{}
			return u, true
		}
	}

	// Associate the current session with the given user name.
	//c.Persist.Info[userSessionKey] = u.Username

	return u, false
}

func (u User) ValidateLogin(v *revel.Validation, user User) {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(user.Password))
	v.Check(err, models.EqualTo{nil}).
		Key("Password").
		Message("Try entering that password again.")
}
