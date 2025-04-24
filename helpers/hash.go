package helpers

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

const tokenCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// GenerateRandomString returns a securely generated random string of length n,
// using a URL-safe base62 charset.
func GenerateRandomString(n int) string {
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		// fallback to empty string on error
		return ""
	}
	for i := range b {
		b[i] = tokenCharset[int(b[i])%len(tokenCharset)]
	}
	return string(b)
}

// HashToken returns the SHA-256 hash of the raw token, as a hex string.
func HashToken(raw string) string {
	sum := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(sum[:])
}
