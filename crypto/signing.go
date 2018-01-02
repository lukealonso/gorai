package crypto

import (
	"bytes"
	"fmt"

	"github.com/lukealonso/ed25519"
)

const PublicKeySize = 32
const PrivateKeySize = 64
const SignatureSize = 64

type PublicKey []byte
type PrivateKey []byte
type Signature []byte

func GenerateKeyPair(hash []byte) (PublicKey, PrivateKey, error) {
	if len(hash) != PublicKeySize {
		return nil, nil, fmt.Errorf("invalid entropy size")
	}
	pub, prv, err := ed25519.GenerateKey(bytes.NewReader(hash), newBlake2b())
	return PublicKey(pub), PrivateKey(prv), err
}

func ExtractPublicKey(privateKey PrivateKey) (PublicKey, error) {
	if len(privateKey) != PrivateKeySize {
		return nil, fmt.Errorf("invalid private key size")
	}
	return PublicKey(ed25519.PrivateKey(privateKey).Public().(ed25519.PublicKey)), nil
}

func VerifySignature(publicKey PublicKey, message []byte, signature Signature) bool {
	if len(signature) != SignatureSize || len(publicKey) != PublicKeySize {
		return false
	}
	return ed25519.Verify(ed25519.PublicKey(publicKey), message, signature, newBlake2b())
}

func Sign(privateKey PrivateKey, message []byte, signature Signature) error {
	if len(signature) != SignatureSize {
		return fmt.Errorf("invalid signature size")
	}
	if len(privateKey) != PrivateKeySize {
		return fmt.Errorf("invalid private key size")
	}
	sig := ed25519.Sign(ed25519.PrivateKey(privateKey), message, newBlake2b())
	if sig == nil || len(sig) != len(signature) {
		return fmt.Errorf("invalid signature size")
	}
	copy(signature, sig)
	return nil
}
