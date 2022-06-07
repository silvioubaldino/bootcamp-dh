package aula2_test

import (
	".bootcamp-dh/Testes/testes1/aula2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSub(t *testing.T) {
	a := 15
	b := 5
	expected := 10

	result := aula2.Sub(a, b)

	assert.Equal(t, expected, result, "Sub falhou")
}

func TestDivide(t *testing.T) {
	a := 15
	b := 5
	expected := 3

	result, _ := aula2.Divide(a, b)

	assert.Equal(t, expected, result, "Sub falhou")
}

func TestDivideFail(t *testing.T) {
	a := 15
	b := 0

	_, err := aula2.Divide(a, b)

	assert.Errorf(t, err, "falhou")
}
