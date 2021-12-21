package core

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/stretchr/testify/assert"
)

func assertTemplateHash(assert *assert.Assertions, template string, expectedHash string) {
	hash := sha1.New()
	hash.Write([]byte(template))
	assert.Equal(expectedHash, hex.EncodeToString(hash.Sum(nil)))
}
