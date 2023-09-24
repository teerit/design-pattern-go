package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPerson(t *testing.T) {
	p := NewPerson("Test", 23)
	assert.Equal(t, "Test", p.Name, "Expecting result to be Test")
	assert.Equal(t, 23, p.Age, "Expecting result to be 23")
}
