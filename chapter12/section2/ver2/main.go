package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"
	"strings"
)

func main() {
	idToken := `[yourIDToken]`

	dataArray := strings.Split(idToken, ".")
	header, payload, sig := dataArray[0], dataArray[1], dataArray[2]

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

	// 公開鍵を使えるようにする
	N := "nnnnnnn"
	E := "eeeeeee"

	dn, _ := base64.RawURLEncoding.DecodeString(N)
	de, _ := base64.RawURLEncoding.DecodeString(E)

	pk := &rsa.PublicKey{
		N: new(big.Int).SetBytes(dn),
		E: int(new(big.Int).SetBytes(de).Int64()),
	}

	// 検証するデータ
	message := sha256.Sum256([]byte(header + "." + payload))

	// 署名をbase64 decodeする
	sigData, err := base64.RawURLEncoding.DecodeString(sig)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	if err := rsa.VerifyPKCS1v15(pk, crypto.SHA256, message[:], sigData); err != nil {
		fmt.Println("invalid token")
	} else {
		fmt.Println("valid token")
		fmt.Println("header: ", string(headerData))
		fmt.Println("payload: ", string(payloadData))
	}
}
