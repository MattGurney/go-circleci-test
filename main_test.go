package main

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestAddition(t *testing.T) {
    assert.Equal(t, 2, 1+1, "1+1 should equal 2")
}

func TestFailure(t *testing.T) {
    assert.Equal(t, 3, 1+1, "This test should fail")
}
