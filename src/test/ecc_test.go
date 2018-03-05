package test

import (
	"strings"
	"testing"

	"../gochain"
)

func TestKeyGeneration(t *testing.T) {

	priv, pub := gochain.GenerateKeys()

	t.Log("ECC Keys: ", priv, pub)

	if len(priv) == 0 {
		t.Error("Private key is empty")
	}

	if len(pub) == 0 {
		t.Error("Private key is empty")
	}

	publicKey := gochain.ParsePublicKey(pub)

	t.Log("Decoded Public Key: (", publicKey.X.String(), ",", publicKey.Y.String(), ")")

	if !strings.Contains(pub, publicKey.X.String()) {
		t.Error("Key encoding problem with X")
	}

	if !strings.Contains(pub, publicKey.Y.String()) {
		t.Error("Key encoding problem with Y")
	}

}
