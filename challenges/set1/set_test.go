package set1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertHexToBase64(t *testing.T) {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	expectedOutput := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	actualOutput, err := ConvertHexToBase64(input)
	if err != nil {
		t.Fatal("")
	}

	assert.Equal(t, expectedOutput, actualOutput)
}
