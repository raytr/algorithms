package word_search

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	board1 = [][]byte{
		{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'},
	}
)

func TestExist(t *testing.T) {
	testCases := []struct {
		name        string
		board       [][]byte
		word        string
		expectation bool
	}{
		{"board = [[\"A\",\"B\",\"C\",\"E\"},{\"S\",\"F\",\"C\",\"S\"},{\"A\",\"D\",\"E\",\"E\"]], word = \"ABCCED\"", board1, "ABCCED", true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expectation, exist(tt.board, tt.word))
		})
	}
}
