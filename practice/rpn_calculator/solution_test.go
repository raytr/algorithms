package main

import (
	"testing"
)

var validTestData = map[string]float64{
	//"-62 -52 ^ 8 -41 in":                    1,
	"-100 not 39 + 84 -":   55.0,  // - -100 + 39 - 84 = 55
	"-63 92 + -31 + -60 -": 58.0,  // -62 + 92 -31 - -60 = 58
	"-31 -92 - 5 * -11 +":  294.0, // (-31 - -92) * 5 -11 = 294
	//"-32.734 not 45.261 -52.036 in 26.69 *": 26.69,
}

var invalidTestData = []string{
	"68 -26 * not + -83 ^",
	"91 38 + ^ -45 -",
	"-19 not * -68 + not",
	"56 -24 97 in -83 ^ 92",
}

func TestValidExpressions(t *testing.T) {
	for expression, expectedResult := range validTestData {
		t.Run(expression, func(t *testing.T) {
			result, err := Calculate(expression)

			if err != nil {
				t.Fatalf("Returned error \"%v\" for valid expression \"%s\".", err, expression)
			}

			if result != expectedResult {
				t.Fatalf("Returned %v for expression \"%v\". Should be %v.", result, expression, expectedResult)
			}
		})
	}
}

func TestInvalidExpressions(t *testing.T) {
	for _, expression := range invalidTestData {
		t.Run(expression, func(t *testing.T) {
			result, err := Calculate(expression)

			if err == nil {
				t.Fatalf("Returned %v for invalid expression \"%s\".", result, expression)
			}
		})
	}
}
