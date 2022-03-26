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
	input1HexDecoded, _ := hexDecodeBytes(input1)
	input2HexDecoded, _ := hexDecodeBytes(input2)

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
