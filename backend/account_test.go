package main

import (
	"fmt"

	"testing"

	"github.com/rayone121/reBank/backend/types"
	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	acc, err := types.NewAccount("John", "Doe", "johndoe", "password")
	assert.Nil(t, err)
	fmt.Printf("Debug info: %+v\n", acc)
}
