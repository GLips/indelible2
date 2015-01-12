package models

import (
	"reflect"
)

type EqualTo struct {
	Expected interface{}
}

func (e EqualTo) IsSatisfied(actual interface{}) bool {
	return reflect.DeepEqual(e.Expected, actual)
}

func (e EqualTo) DefaultMessage() string {
	return "Things are not equal!"
}
