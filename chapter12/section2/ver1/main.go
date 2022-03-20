package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {
	idToken := `[yourIDToken]`

	dataArray := strings.Split(idToken, ".")
	header, payload, _ := dataArray[0], dataArray[1], dataArray[2]

	// headerをbase64 decodeする
	headerData, err := base64.RawURLEncoding.DecodeString(header)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// payloadをbase64 decodeする
	payloadData, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Println("header: ", string(headerData))
	fmt.Println("payload: ", string(payloadData))
}
