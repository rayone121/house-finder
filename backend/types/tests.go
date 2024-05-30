package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	acc, err := NewAccount("John", "Doe", "johndoe", "password")
	assert.Nil(t, err)
	fmt.Printf("Debug info: %+v\n", acc)
}
