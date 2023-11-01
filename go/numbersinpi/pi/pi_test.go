package pi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPI_NumbersInPi(t *testing.T) {
	tests := []struct {
		name     string
		scenario func(t *testing.T)
	}{
		{
			name: "should find the number of piece to form the pi number",
			scenario: func(t *testing.T) {
				const PI = "3141592653589793238462643383279"
				numbers := []string{"314159265358979323846", "26433", "8", "3279", "314159265", "35897932384626433832", "79"}
				assert.Equal(t, 2, NumbersInPi(PI, numbers))

			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.scenario(t)
		})
	}
}
