package unid

import (
	"crypto/rand"
	"fmt"
	"log"
)

/* fake change */

type unid []byte

var uindCounter = get64BitRandomNumber()
var uintProcessUnique = readRandomBytes(6)

func get64BitRandomNumber() uint64 {

	var rand_number uint64
	rand_bytes := make([]byte, 8)

	_, err := rand.Read(rand_bytes)

	if err != nil {
		panic(fmt.Errorf("failed to read random bytes"))
	}

	rand_number = (uint64(rand_bytes[0]) << 56) |
		(uint64(rand_bytes[1]) << 48) |
		(uint64(rand_bytes[2]) << 40) |
		(uint64(rand_bytes[3]) << 32) |
		(uint64(rand_bytes[4]) << 24) |
		(uint64(rand_bytes[5]) << 16) |
		(uint64(rand_bytes[6]) << 8) |
		(uint64(rand_bytes[7]) << 0)

	return rand_number
}

func readRandomBytes(n_bytes int) []byte {

	rand_bytes := make([]byte, n_bytes)

	_, err := rand.Read(rand_bytes)

	if err != nil {
		panic(fmt.Errorf("failed to read random bytes"))
	}
	return rand_bytes
}

func getUindCounterValue() uint64 {
	ret := uindCounter
	uindCounter++
	return ret
}

func GetUNID() unid {
	var b []byte

	log.Println(uintProcessUnique)
	log.Println(getUindCounterValue())

	return b
}
