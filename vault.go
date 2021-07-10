package vault

import (
	"errors"
	"github.com/YMajid/vault/encrypt"
)

type Vault struct {
	encodingKey string
	keyValues   map[string]string
}

func MemoryVault(encodingKey string) Vault {
	return Vault{encodingKey: encodingKey, keyValues: make(map[string]string)}
}

func (v *Vault) Get(key string) (string, error) {
	hex, ok := v.keyValues[key]
	if !ok {
		return "", errors.New("vault: no value for that key")
	}

	ret, err := encrypt.Decrypt(v.encodingKey, hex)
	if err != nil {
		return "", err
	}

	return ret, nil
}

func (v *Vault) Set(key string, value string) error {
	encryptedValue, err := encrypt.Encrypt(v.encodingKey, value)
	if err != nil {
		return err
	}

	v.keyValues[key] = encryptedValue
	return nil
}
