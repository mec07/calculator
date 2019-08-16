package calculator_test

import (
	"testing"

	"github.com/mec07/calculator"
	"github.com/stretchr/testify/assert"
)

func TestErrorCase(t *testing.T) {
	assert.Equal(t, "2+a", calculator.Calculate("2+a"))
}

func TestSimpleAddition(t *testing.T) {
	assert.Equal(t, "4", calculator.Calculate("2+2"))
}
