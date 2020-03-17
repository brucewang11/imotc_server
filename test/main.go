package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	var salt = "wangbruce"
	sha := sha256.New()
	sha.Write([]byte("runchain666"))
	sha.Write([]byte(salt))
	pass := sha.Sum(nil)
	aa := hex.EncodeToString(pass)
	fmt.Println(aa)
}
