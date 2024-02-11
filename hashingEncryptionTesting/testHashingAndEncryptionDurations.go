package main

import (
	"CMPSC488SP24SecThursday/hashing"
	"fmt"
)

func main() {
	fmt.Println("MD5 Duration Test:")
	fmt.Println(hashing.TestMD5Duration())
	fmt.Println()

	fmt.Println("SHA1 Duration Test:")
	fmt.Println(hashing.TestSHA1Duration())
	fmt.Println()

	fmt.Println("SHA256 Duration Test:")
	fmt.Println(hashing.TestSHA256Duration())
	fmt.Println()
}
