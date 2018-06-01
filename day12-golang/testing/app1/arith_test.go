package main

import (
    "testing"

    "github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
    total := Add(1, 9)
    assert.NotNil(t, total, "The `total` should not be `nil`")
    assert.Equal(t, 4, total, "Expecting `4`")
}

/*
func TestSubtract(t *testing.T) {
    total := Subtract(1, 9)
    assert.NotNil(t, total, "The `total` should not be `nil`")
    assert.Equal(t, -2, total, "Expecting `-2`")
}
*/
// go test