package har

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParseHar(t *testing.T) {
	harFileContent, err := os.ReadFile("./data/www.google.com.har")
	assert.Nil(t, err)
	har, err := ParseHar(harFileContent)
	assert.Nil(t, err)
	t.Log(har)
}
