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

func randomStringWithPrefix(prefix string) (string, error) {
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

		id := prefix + string(result)
		// Check if ID exists in the map
		if _, exists := generatedIDs[id]; !exists {
			generatedIDs[id] = true
			return id, nil
		}
	}
}

//func main() {
//	// Generate ID with prefix 'g'
//	id, err := randomStringWithPrefix("g")
//	if err != nil {
//		fmt.Println("Error generating random string:", err)
//		return
//	}
//	fmt.Println("Generated ID:", id)
//}

//package main
//
//import (
//"crypto/rand"
//"errors"
//"fmt"
//"math/big"
//)
//
//const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
//const length = 4
//
//var generatedIDs = make(map[string]bool)

func incrementID(id string) string {
	base := len(charset)
	bytes := []byte(id)

	for i := len(bytes) - 1; i >= 0; i-- {
		index := charsetIndexOf(bytes[i])
		if index+1 < base {
			bytes[i] = charset[index+1]
			break
		} else {
			bytes[i] = charset[0]
		}
	}

	return string(bytes)
}

func charsetIndexOf(char byte) int {
	for i, c := range charset {
		if byte(c) == char {
			return i
		}
	}
	return -1
}

func randomStringWithPrefixMk2(prefix string) (string, error) {
	result := make([]byte, length)
	charsetLen := big.NewInt(int64(len(charset)))

	for i := range result {
		randomInt, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			return "", err
		}
		result[i] = charset[randomInt.Int64()]
	}

	id := prefix + string(result)

	for {
		if _, exists := generatedIDs[id]; !exists {
			generatedIDs[id] = true
			return id, nil
		}

		id = prefix + incrementID(id[len(prefix):])

		// Check for exhaustion
		if id[len(id)-length:] == charset[:length] {
			return "", errors.New("ID space exhausted")
		}
	}
}

//func main() {
//	// Generate ID with prefix 'g'
//	id, err := randomStringWithPrefix("g")
//	if err != nil {
//		fmt.Println("Error generating random string:", err)
//		return
//	}
//	fmt.Println("Generated ID:", id)
//}
