package middleware

import (
	"crypto/rand"
	"math/big"
)

//package main
//
//import (
//"crypto/rand"
//"fmt"
//"math/big"
//)

const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
const length = 4

var generatedIDs = make(map[string]bool)

func randomStringWithPrefix() (string, error) {
	var err error
	for {
		result := make([]byte, length)
		charsetLen := big.NewInt(int64(len(charset)))

		for i := range result {
			randomInt, err := rand.Int(rand.Reader, charsetLen)
			if err != nil {
				return "", err
			}
			result[i] = charset[randomInt.Int64()]
		}

		id := string(result)
		// Check if ID exists in the map
		if _, exists := generatedIDs[id]; !exists {
			generatedIDs[id] = true
			return id, err
		} else {
			id, err = randomStringWithPrefix()
			return id, err
		}
	}
}
