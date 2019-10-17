package random

import (
	"encoding/base64"
	"fmt"
)

func NewTokenGenerator(length uint32) TokenGenerator {
	return &token{length: length}
}

type token struct {
	length uint32
}

func (t *token) Generate() (string, error) {
	b, err := generateRandomBytes(t.length)
	if err != nil {
		return "", fmt.Errorf("generating random bytes error: %w", err)
	}

	return base64.URLEncoding.EncodeToString(b)[:t.length], nil
}
