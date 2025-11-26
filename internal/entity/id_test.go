package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIDTypeValue(t *testing.T) {
	assert.Equal(t, ID("account:123").Type(), "account")
	assert.Equal(t, ID("account:123").ID(), "123")
}
