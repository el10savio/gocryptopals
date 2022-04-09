package set1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertHexToBase64(t *testing.T) {
	input := []byte("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	expectedOutput := []byte("SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t")

	actualOutput, err := ConvertHexToBase64(input)
	if err != nil {
		t.Fatal("error received from ConvertHexToBase64:", err)
	}

	assert.Equal(t, expectedOutput, actualOutput)
}

func TestFixedXOR(t *testing.T) {
	input1 := []byte("1c0111001f010100061a024b53535009181c")
	input2 := []byte("686974207468652062756c6c277320657965")
	expectedOutput := []byte("746865206b696420646f6e277420706c6179")

	actualOutput, err := FixedXOR(input1, input2)
	if err != nil {
		t.Fatal("error received from FixedXOR:", err)
	}

	assert.Equal(t, expectedOutput, actualOutput)
}

func TestSingleByteXORCipher(t *testing.T) {
	input := []byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	expectedOutput := []byte("Cooking MC's like a pound of bacon")
	expectedScore := 204

	actualOutput, actualScore, err := SingleByteXORCipher(input)
	if err != nil {
		t.Fatal("error received from SingleByteXORCipher:", err)
	}

	assert.Equal(t, expectedOutput, actualOutput)
	assert.Equal(t, expectedScore, actualScore)
}

func TestDetectSingleCharacterXOR(t *testing.T) {
	input := "./files/4.txt"
	expectedOutput := []byte("Now that the party is jumping\n")

	actualOutput, err := DetectSingleCharacterXOR(input)
	if err != nil {
		t.Fatal("error received from DetectSingleCharacterXOR:", err)
	}

	assert.Equal(t, expectedOutput, actualOutput)
}
