package main

import (
	// "bytes"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func passCrack(password string) string {

	salt := password[:2]
	var temp string
	var temp2 string
	fmt.Println("Salt is", salt)
	fmt.Println("Hash is", password[2:])
	for x := 0; x < 10000; x++ {
		h := sha1.New()
		temp = salt + strconv.Itoa(x)
		fmt.Println(x)
		io.WriteString(h, temp)
		temp3 := h.Sum(nil)
		temp2 = strings.ToUpper(hex.EncodeToString(temp3))
		// fmt.Println(temp2)
		if temp2 == password[2:] {
			fmt.Println("The password is ", x)
			return temp
		}
		// fmt.Printf("% x", h.Sum(nil))
	}
	return ""
}

// TO RUN -->
// go build q1.go
// ./q1 <hash>
// pin is 1701
func main() {
	argsWithoutProg := os.Args[1]
	passCrack(argsWithoutProg)
}
