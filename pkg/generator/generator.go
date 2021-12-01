package generator

import (
	"math/rand"
	"strings"

	"github.com/gofrs/uuid"
)

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits

	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	defaultIDLength = 38
)

type UUIDGenerator interface {
	GenerateString() string
}

type uuidGenerator struct {
	gen uuid.Generator
}

func (u *uuidGenerator) GenerateString() string {
	uuidVal, err := u.gen.NewV4()
	if err != nil {
		return randStringBytesMaskImprSrcSB(defaultIDLength)
	}

	return uuidVal.String()
}

func NewUUIDGenerator(gen uuid.Generator) UUIDGenerator {
	return &uuidGenerator{gen: gen}
}

func randStringBytesMaskImprSrcSB(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}
