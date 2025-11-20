package entity

import (
	"fmt"
	"strings"
)

type ID string

func (id ID) Type() string {
	return strings.Split(string(id), ":")[0]
}

func (id ID) ID() string {
	return strings.Split(string(id), ":")[1]
}

func (id ID) String() string {
	return string(id)
}

func NewID[T any](typ string, id T) ID {
	return ID(fmt.Sprintf("%s:%v", typ, id))
}