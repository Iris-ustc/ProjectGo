package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"io"
	"strings"
)

func main() {
	
    sHeader := "this is header."
    sBody := "this is body."
	sParam := fmt.Sprintf(`{"Head":%s,"Body":%s}`, sHeader, sBody)

	sKey := "this is key"
	sAferEnParam := AesEncryptECB(sParam, sKey)
	fmt.Println("AES En:\n", sAferEnParam)
	
	sAferDeParam := AesDecryptECB(sAferEnParam, sKey)
	fmt.Println("AES De:\n", sAferDeParam)
	
}
