package aula2_test

import (
	".bootcamp-dh/Testes/testes1/aula2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	var list []int
	list = append(list, 3, 2, 5, 4, 6)
	var expected []int
	expected = append(expected, 2, 3, 4, 5, 6)

	ordered := aula2.Sort(list)

	assert.Equal(t, expected, ordered)
}
