package configuration

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestReadingWritingConfig(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "")
	filename := Filename(tmpDir)
	assert.NoError(t, err)

	writtenConfig := Config{
		GithubUsername:    "user",
		ApiKey:            "MyKey",
		ExercismDirectory: "/exercism/directory",
	}

	ToFile(filename, writtenConfig)

	loadedConfig, err := FromFile(filename)
	assert.NoError(t, err)

	assert.Equal(t, writtenConfig, loadedConfig)
}
