package devto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProvider(t *testing.T) {
	provider := Provider()
	assert.NoError(t, provider.InternalValidate())
}
