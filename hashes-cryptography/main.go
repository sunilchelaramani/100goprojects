package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func encodingMD5(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
	var password string

	fmt.Print("Enter password: ")
	fmt.Scan(&password)

	fmt.Println("MD5: ", encodingMD5(password))
}
