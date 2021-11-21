package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	p := Passport{
		BirthYear: "2021",
	}

	fmt.Println(validate(p))
	assert.Error(t, nil)
}
