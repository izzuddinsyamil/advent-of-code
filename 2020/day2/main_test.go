package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessInput(t *testing.T) {
	tests := map[string]struct {
		input inputLine
		want  bool
	}{
		"valid 1": {
			input: inputLine{
				pos1:          0,
				pos2:          2,
				charToInclude: "a",
				password:      "abcde",
			},
			want: true,
		},
		"valid 2": {
			input: inputLine{
				pos1:          4,
				pos2:          5,
				charToInclude: "d",
				password:      "dddtpdd",
			},
			want: true,
		},
		"invalid 1": {
			input: inputLine{
				pos1:          0,
				pos2:          2,
				charToInclude: "b",
				password:      "cdefg",
			},
			want: false,
		},
		"invalid 2": {
			input: inputLine{
				pos1:          1,
				pos2:          8,
				charToInclude: "c",
				password:      "ccccccccc",
			},
			want: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := processInput(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}
}
