package generator

import (
	"testing"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGenerateString(t *testing.T) {
	gen := NewUUIDGenerator(uuid.NewGen())
	assert.NotEmpty(t, gen.GenerateString())
	assert.NotEmpty(t, randStringBytesMaskImprSrcSB(defaultIDLength))
}
