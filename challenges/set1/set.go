package set1

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
)

// Challenge 1: Convert hex to base64
func ConvertHexToBase64(input string) (string, error) {
	hexBytes := []byte(input)

	hexDecoded := make([]byte, hex.DecodedLen(len(hexBytes)))
	_, err := hex.Decode(hexDecoded, hexBytes)
	if err != nil {
		return "", err
	}

	// fmt.Println(string(hexDecoded))

	base64Encoded := make([]byte, base64.StdEncoding.EncodedLen(len(input)))
	base64.StdEncoding.Encode(base64Encoded, hexDecoded)
	base64Encoded = bytes.Trim(base64Encoded, "\x00")

	return string(base64Encoded), nil
}
