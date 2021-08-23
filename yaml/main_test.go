package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	err := app.Run([]string{"-o", "./test.yaml"})
	assert.NoError(t, err)
}
