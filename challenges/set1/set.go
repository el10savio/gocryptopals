package set1

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"errors"
)

// Challenge 1: Convert hex to base64
func ConvertHexToBase64(hexBytes []byte) ([]byte, error) {
	hexDecoded, err := hexDecodeBytes(hexBytes)
	if err != nil {
		return []byte{}, err
	}

	// fmt.Println(string(hexDecoded))

	base64Encoded := base64EncodeBytes(hexDecoded)
	base64Encoded = bytes.Trim(base64Encoded, "\x00")

	return base64Encoded, nil
}

// Challenge 2: FixedXOR
func FixedXOR(input1, input2 []byte) ([]byte, error) {
	input1HexDecoded, err := hexDecodeBytes(input1)
	if err != nil {
		return []byte{}, err
	}

	input2HexDecoded, err := hexDecodeBytes(input2)
	if err != nil {
		return []byte{}, err
	}

	if len(input1HexDecoded) != len(input2HexDecoded) {
		return []byte{}, errors.New("length of inputs in bytes are not equal")
	}

	outputBytes := make([]byte, len(input1HexDecoded))
	for index := 0; index < len(input1HexDecoded); index++ {
		outputBytes[index] = input1HexDecoded[index] ^ input2HexDecoded[index]
	}

	// fmt.Println(string(outputBytes))

	outputHexEncodedBytes := hexEncodeBytes(outputBytes)

	return outputHexEncodedBytes, nil
}

// Challenge 3: Single-byte XOR cipher
func SingleByteXORCipher(input []byte) ([]byte, int, error) {
	inputHexDecoded, err := hexDecodeBytes(input)
	if err != nil {
		return []byte{}, 0, err
	}

	var result []byte
	scoreMax := 0

	for i := 0; i < 256; i++ {
		resultPossible := make([]byte, len(inputHexDecoded))
		scoreSoFar := 0

		for j := 0; j < len(inputHexDecoded); j++ {
			char := inputHexDecoded[j] ^ byte(i)
			scoreSoFar += weightMap(char)
			resultPossible[j] = char
		}

		if scoreSoFar > scoreMax {
			result, scoreMax = resultPossible, scoreSoFar
		}

		scoreSoFar = 0
	}

	return result, scoreMax, nil
}

func hexDecodeBytes(hexBytes []byte) ([]byte, error) {
	hexDecoded := make([]byte, hex.DecodedLen(len(hexBytes)))
	_, err := hex.Decode(hexDecoded, hexBytes)
	return hexDecoded, err
}

func hexEncodeBytes(input []byte) []byte {
	hexEncoded := make([]byte, hex.EncodedLen(len(input)))
	_ = hex.Encode(hexEncoded, input)
	return hexEncoded
}

func base64DecodeBytes(base64Bytes []byte) []byte {
	base64Decoded := make([]byte, base64.StdEncoding.DecodedLen(len(base64Bytes)))
	base64.StdEncoding.Encode(base64Decoded, base64Bytes)
	return base64Decoded
}

func base64EncodeBytes(input []byte) []byte {
	base64Encoded := make([]byte, base64.StdEncoding.EncodedLen(len(input)))
	base64.StdEncoding.Encode(base64Encoded, input)
	return base64Encoded
}

func weightMap(char byte) int {
	weight := map[byte]int{
		byte('U'): 2,
		byte('u'): 2,
		byte('L'): 3,
		byte('l'): 3,
		byte('D'): 4,
		byte('d'): 4,
		byte('R'): 5,
		byte('r'): 5,
		byte('H'): 6,
		byte('h'): 6,
		byte('S'): 7,
		byte('s'): 7,
		byte(' '): 8,
		byte('N'): 9,
		byte('n'): 9,
		byte('I'): 10,
		byte('i'): 10,
		byte('O'): 11,
		byte('o'): 11,
		byte('A'): 12,
		byte('a'): 12,
		byte('T'): 13,
		byte('t'): 13,
		byte('E'): 14,
		byte('e'): 14,
	}
	return weight[char]
}
