package gochain

import (
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"math/big"
	"strings"
)

type PublicKey struct {
	X big.Int
	Y big.Int
}

func GenerateKeys() (string, string) {

	priv, x, y, _ := elliptic.GenerateKey(elliptic.P256(), rand.Reader)

	return hex.EncodeToString(priv), EncodePublicKey(x, y)
}

func EncodePublicKey(x, y *big.Int) string {
	return x.String() + "|" + y.String()
}

func ParsePublicKey(encodedKey string) PublicKey {

	tokens := strings.Split(encodedKey, "|")

	var key PublicKey
	key.X.SetString(tokens[0], 10)
	key.Y.SetString(tokens[1], 10)

	return key

}
