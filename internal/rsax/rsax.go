package rsax

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"io/ioutil"

	"github.com/james-lawrence/pacmir/internal/fsx"
	"github.com/pkg/errors"
)

const (
	defaultBits = 4096 // 4096 bit keysize, reasonable default.
)

// Auto generates a RSA key using package defined defaults.
func Auto() (pkey []byte, err error) {
	return Generate(defaultBits)
}

// CachedAuto loads/generates an RSA key at the provided filepath.
func CachedAuto(path string) (pkey []byte, err error) {
	if fsx.FileExists(path) {
		return ioutil.ReadFile(path)
	}

	if pkey, err = Auto(); err != nil {
		return nil, err
	}

	if err = ioutil.WriteFile(path, pkey, 0600); err != nil {
		return nil, err
	}

	return pkey, nil
}

// CachedGenerate loads/generates an SSH key at the provided filepath.
func CachedGenerate(path string, bits int) (pkey []byte, err error) {
	if fsx.FileExists(path) {
		return ioutil.ReadFile(path)
	}

	if pkey, err = Generate(bits); err != nil {
		return nil, err
	}

	if err = ioutil.WriteFile(path, pkey, 0600); err != nil {
		return nil, err
	}

	return pkey, nil
}

// UnsafeAuto generates a ssh key using unsafe defaults, this method is used to
// generate an ssh key quickly for cases were we do not care about safety, i.e.
// tests.
func UnsafeAuto() (pkey []byte, err error) {
	return Generate(128)
}

// Generate a RSA private key with the given bits size, returns the pem encoded bytes.
func Generate(bits int) (encoded []byte, err error) {
	var (
		pkey *rsa.PrivateKey
	)

	if pkey, err = private(bits); err != nil {
		return encoded, err
	}

	// Get ASN.1 DER format
	marshalled := x509.MarshalPKCS1PrivateKey(pkey)

	return pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: marshalled,
	}), nil
}

// generatePrivateKey creates a RSA Private Key of specified byte size
func private(bits int) (k *rsa.PrivateKey, err error) {
	// Private Key generation
	if k, err = rsa.GenerateKey(rand.Reader, bits); err != nil {
		return k, err
	}

	// Validate Private Key
	return k, k.Validate()
}

// PublicKey returns a public key from the pem encoded private key.
func PublicKey(pemkey []byte) (pub []byte, err error) {
	var (
		pkey *rsa.PrivateKey
	)

	blk, _ := pem.Decode(pemkey) // assumes a single valid pem encoded key.

	if pkey, err = x509.ParsePKCS1PrivateKey(blk.Bytes); err != nil {
		return pub, err
	}

	return x509.MarshalPKCS1PublicKey(&pkey.PublicKey), nil
}

// Decode decode a RSA private key.
func Decode(encoded []byte) (priv *rsa.PrivateKey, err error) {
	b, _ := pem.Decode(encoded)
	if priv, err = x509.ParsePKCS1PrivateKey(b.Bytes); err != nil {
		return nil, errors.WithStack(err)
	}

	return priv, nil
}

// DecodePKCS1PrivateKey decode PEM to x509.PKCS1PrivateKey bytes
func DecodePKCS1PrivateKey(encoded []byte) []byte {
	b, _ := pem.Decode(encoded)
	return b.Bytes
}

// MaybeDecode decodes RSA from an encoded array and possible error.
func MaybeDecode(encoded []byte, err error) (priv *rsa.PrivateKey, _ error) {
	if err != nil {
		return priv, err
	}
	return Decode(encoded)
}

// MustDecode decodes RSA from an encoded array and possible error.
func MustDecode(encoded []byte) (priv *rsa.PrivateKey) {
	var err error

	if priv, err = Decode(encoded); err != nil {
		panic(err)
	}

	return priv
}

// FingerprintSHA256 of the key
func FingerprintSHA256(b []byte) string {
	digest := sha256.Sum256(b)
	return hex.EncodeToString(digest[:])
}
