package models

import (
	"time"
)

type Tmpl struct {
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

type Model interface {
	Touch()
	Update()
	Create()
	Delete()

	// One and Many return the keys for this model that are used by Ember
	// in the JSON request/response.
	One() string
	Many() string

	PostProcess()
}

func (t Tmpl) Touch() {
	createdAt := t.CreatedAt
	if &createdAt == nil {
		t.CreatedAt = time.Now()
	}
	t.UpdatedAt = time.Now()
}

func (t Tmpl) PostProcess() {
}
