package hasher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	hashLen = 10

	value = "test_string_here"
)

func TestHasherReverseString(t *testing.T) {
	h := NewHasher(hashLen)
	encoded := h.EncodeString(value)
	assert.NotEmpty(t, encoded)

	decoded, err := h.DecodeString(encoded)
	assert.NoError(t, err)
	assert.Equal(t, value, decoded)
}
