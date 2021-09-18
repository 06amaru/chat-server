package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
)

func createHash(key string) string {
	// New returns a new hash.Hash computing the MD5 checksum.
	hasher := md5.New()
	// Write adds more data to the running hash.
	hasher.Write([]byte(key))

	// Sum appends the current hash to b and returns the resulting slice.
	// It does not change the underlying hash state.
	// Sum(b []byte) []byte
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passhprase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passhprase)))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	//https://www.youtube.com/watch?v=jgTqR8PuWuU
}
