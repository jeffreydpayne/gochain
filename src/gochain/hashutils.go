package gochain

import (
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/ripemd160"
)

func SingleHash(contents string) string {

	h := sha256.New()
	h.Write([]byte(contents))
	return hex.EncodeToString(h.Sum(nil))

}

func DoubleHash(contents string) string {

	first := sha256.New()
	first.Write([]byte(contents))

	second := sha256.New()
	second.Write(first.Sum(nil))

	return hex.EncodeToString(second.Sum(nil))

}

func RIPEMD_160(contents []byte) string {

	h := ripemd160.New()
	h.Write([]byte(contents))
	return hex.EncodeToString(h.Sum(nil))

}

func RIPEMD_160_With_SHA_256(contents string) string {

	first := sha256.New()
	first.Write([]byte(contents))

	second := ripemd160.New()
	second.Write(first.Sum(nil))
	return hex.EncodeToString(second.Sum(nil))

}

func MerkleHash(h1 string, h2 string) string {

	first := sha256.New()

	b1, _ := hex.DecodeString(h1)
	b2, _ := hex.DecodeString(h2)

	first.Write(b1)
	first.Write(b2)

	second := sha256.New()
	second.Write(first.Sum(nil))

	return hex.EncodeToString(second.Sum(nil))

}
