package random

import (
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"
)

var (
	ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)

type argon struct {
	params *params
}

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

func NewHashGenerator() HashGenerator {
	return &argon{
		params: &params{
			memory:      128 * 1024,
			iterations:  4,
			parallelism: 4,
			saltLength:  16,
			keyLength:   32,
		},
	}
}

func (a *argon) Generate(password string) (string, error) {
	salt, err := generateRandomBytes(a.params.saltLength)
	if err != nil {
		return "", fmt.Errorf("generating new password hash err: %w", err)
	}

	hash := argon2.IDKey([]byte(password), salt, a.params.iterations,
		a.params.memory, a.params.parallelism, a.params.keyLength)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Return a string using the standard encoded hash representation.
	encodedHash := "" +
		"$argon2id$v=" + strconv.Itoa(argon2.Version) +
		"$m=" + strconv.Itoa(int(a.params.memory)) +
		",t=" + strconv.Itoa(int(a.params.iterations)) +
		",p=" + strconv.Itoa(int(a.params.parallelism)) +
		"$" + b64Salt + "$" + b64Hash

	return encodedHash, nil
}

func (a *argon) Compare(password, encodedHash string) (bool, error) {

	p, salt, hash, err := a.decodeHash(encodedHash)
	if err != nil {
		return false, fmt.Errorf("decoding hash err: %w", err)
	}

	otherHash := argon2.IDKey([]byte(password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}

func (a *argon) decodeHash(encodedHash string) (*params, []byte, []byte, error) {
	splits := strings.Split(encodedHash, "$")
	if len(splits) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}

	var version int
	_, err := fmt.Sscanf(splits[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}

	p := &params{}
	_, err = fmt.Sscanf(splits[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(splits[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.saltLength = uint32(len(salt))

	hash, err := base64.RawStdEncoding.DecodeString(splits[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.keyLength = uint32(len(hash))

	return p, salt, hash, nil
}
