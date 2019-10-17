package random

import "crypto/rand"

type HashGenerator interface {
	Generate(password string) (string, error)
	Compare(password, encodedHash string) (bool, error)
}

type TokenGenerator interface {
	Generate() (string, error)
}

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
