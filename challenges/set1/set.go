package set1

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
)

// Challenge 1: Convert hex to base64
func ConvertHexToBase64(input string) (string, error) {
	hexBytes := []byte(input)

	hexDecoded, err := hexDecodeBytes(hexBytes)
	if err != nil {
		return "", err
	}

	// fmt.Println(string(hexDecoded))

	base64Encoded := base64EncodeBytes(hexDecoded)
	base64Encoded = bytes.Trim(base64Encoded, "\x00")

	return string(base64Encoded), nil
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
