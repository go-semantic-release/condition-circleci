package condition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGithubValid(t *testing.T) {
	cci := CircleCI{}
	err := cci.RunCondition(map[string]string{})
	assert.NoError(t, err)
}
