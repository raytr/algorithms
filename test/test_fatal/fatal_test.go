package test_fatal

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFatal(t *testing.T) {
	testCase := []struct {
		name  string
		value int
	}{
		{"case 1", 1},
		{"case 2", 2},
	}

	for _, tt := range testCase {
		t.Fatal(tt.name, func(t *testing.T) {
			assert.Equal(t, 0, PrintNumber(tt.name, tt.value))
		})

		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, 0, PrintNumber(tt.name, tt.value))
		})
	}
}

func PrintNumber(s string, i int) int {
	fmt.Println(fmt.Sprintf("name: %v, val %v"), s, i)
	log.Fatal()
	return 0
}
