package test

import (
	"strings"
	"testing"

	"../gochain"
)

func TestSingleHash(t *testing.T) {

	hashedValue := gochain.SingleHash("The quick brown fox jumped over the lazy dog.")

	t.Log("SHA156(X): " + hashedValue)

	if !strings.EqualFold(hashedValue, "68B1282B91DE2C054C36629CB8DD447F12F096D3E3C587978DC2248444633483") {
		t.Error("Wrong hash value.")
	}

}

func TestDoubleHash(t *testing.T) {

	hashedValue := gochain.DoubleHash("The quick brown fox jumped over the lazy dog.")

	t.Log("SHA156(SHA256(X)): " + hashedValue)

	if strings.EqualFold(hashedValue, "68B1282B91DE2C054C36629CB8DD447F12F096D3E3C587978DC2248444633483") {
		t.Error("Single hash returned.")
	}

}

func TestSingleRipemd160(t *testing.T) {

	hashedValue := gochain.RIPEMD_160([]byte("The quick brown fox jumped over the lazy dog."))

	t.Log("RIPEMD160(X): " + hashedValue)

	if !strings.EqualFold(hashedValue, "ec457d0a974c48d5685a7efa03d137dc8bbde7e3") {
		t.Error("Wrong hash value.")
	}

}

func TestDoubleRipemd160(t *testing.T) {

	hashedValue := gochain.RIPEMD_160_With_SHA_256("The quick brown fox jumped over the lazy dog.")

	t.Log("RIPEMD160(SHA256(X)): " + hashedValue)

	if strings.EqualFold(hashedValue, "ec457d0a974c48d5685a7efa03d137dc8bbde7e3") {
		t.Error("Wrong hash value.")
	}

}

func TestMerkleHash(t *testing.T) {

	h1 := gochain.DoubleHash("The quick brown fox ")
	h2 := gochain.DoubleHash("jumped over the lazy dog")

	merkleHash := gochain.MerkleHash(h1, h2)

	t.Log("Merkle Hash: " + merkleHash)

	if len(merkleHash) == 0 {
		t.Error("Empty merkle hash.")
	}

}
