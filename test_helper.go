package core

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/stretchr/testify/assert"
	config "github.com/tommzn/go-config"
)

func loadConfigForTest(fileName *string) config.Config {

	configFile := "fixtures/testconfig.yml"
	if fileName != nil {
		configFile = *fileName
	}
	configLoader := config.NewFileConfigSource(&configFile)
	config, _ := configLoader.Load()
	return config
}

func assertTemplateHash(assert *assert.Assertions, template string, expectedHash string) {
	hash := sha1.New()
	hash.Write([]byte(template))
	assert.Equal(expectedHash, hex.EncodeToString(hash.Sum(nil)))
}
