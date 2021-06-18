package main

import (
	"NetLikePlate/controllers"
	_ "NetLikePlate/controllers"
	_ "NetLikePlate/properties"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSomething(t *testing.T) {
	x := controllers.Testing123(1, 2)
	assert.Equal(t,x,3, "equal")
	assert.NotEqual(t,x,5, "not equal")
}
