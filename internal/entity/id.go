package entity

import "strings"

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