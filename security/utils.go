package security

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{};:,.<>/?"
const length = 30

func GenerateKey() []byte {
	key := make([]byte, length)
	charsetLength := big.NewInt(int64(len(charset)))

	for i := 0; i < 30; i++ {
		index, _ := rand.Int(rand.Reader, charsetLength)
		key[i] = charset[index.Int64()]
	}

	return key
}

func CreateHash(key string) string {
	// New returns a new hash.Hash computing the MD5 checksum.
	hasher := md5.New()
	// Write adds more data to the running hash.
	hasher.Write([]byte(key))

	// Sum appends the current hash to b and returns the resulting slice.
	// It does not change the underlying hash state.
	// Sum(b []byte) []byte
	return hex.EncodeToString(hasher.Sum(nil))
}

func Encrypt(data []byte, passhprase string) []byte {
	key := []byte(CreateHash(passhprase))

	block, _ := aes.NewCipher(key)

	// AEAD (authenticated encryption with associated data) is a cipher mode
	// providing authenticated encryption with associated data.
	gcm, _ := cipher.NewGCM(block)

	// nonce is a kind of initialization vector
	// declare size of nonce
	nonceSize := gcm.NonceSize()
	// allocates nonce
	nonce := make([]byte, nonceSize)
	// fill nonce with random values
	io.ReadFull(rand.Reader, nonce)

	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.
	//
	// To reuse plaintext's storage for the encrypted output, use plaintext[:0]
	// as dst. Otherwise, the remaining capacity of dst must not overlap plaintext.
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func Decrypt(data []byte, passhprase string) []byte {
	key := []byte(CreateHash(passhprase))

	block, _ := aes.NewCipher(key)

	gcm, _ := cipher.NewGCM(block)

	nonceSize := gcm.NonceSize()

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plaintext, _ := gcm.Open(nil, nonce, ciphertext, nil)

	return plaintext
}
