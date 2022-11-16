package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/rand"
	"crypto/x509"
	"crypto/sha512"
)



// Generate keys that can be used immediately.
//
func GeneratePubPriKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	priKey, err := rsa.GenerateKey(rand.Reader, 4096) 
	if err != nil {
		return nil, nil, err
	}
	// The public key is a part of the *rsa.PrivateKey struct
	pubKey := &priKey.PublicKey

	return priKey, pubKey, nil
}

func SignTxt(plainTxt []byte, privateKey *rsa.PrivateKey) ([]byte, error) {
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto

	// Hash the text before signing. Will sign the hashed value.
	newHash := sha512.New()
	if _, err := newHash.Write(plainTxt); err != nil {
		return nil, err
	}
	hashedTxt := newHash.Sum(nil)

	// To generate a signature, need to provide 
	//   - a random number generator
	//   - private key
	//   - hashing algorithm used to hash the plaintext
	//   - hash sum of the plaintext
	signature, err := rsa.SignPSS(
		rand.Reader,
		privateKey,
		crypto.SHA512,
		hashedTxt,
		&opts,
	)
	if err != nil {
		return nil, err
	}

	return signature, nil
}

func VerifyTxt(plainTxt []byte, publicKey *rsa.PublicKey, signature []byte) error {
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto

	// Hashed the text before signing. So need to do the same to get the hashed text.
	newHash := sha512.New()
	if _, err := newHash.Write(plainTxt); err != nil {
		return err
	}
	hashedTxt := newHash.Sum(nil)

	// To verify the signature, need to provide 
	//   - public key
	//   - hashing algorithm used to hash the plaintext (same in signing)
	//   - hash sum of the plaintext (same in signing process)
	//   - the signature
	//   - optional "options"
	err := rsa.VerifyPSS(
		publicKey,
		crypto.SHA512,
		hashedTxt,
		signature,
		&opts,
	)
	if err != nil {
		return err
	}

	// if no error, means the signature is valid, 
	// verification success.
	return nil
}


// Generate private/public keys in bytes
func GeneratePubPriKeyPairBytes() ([]byte, []byte, error) {
	priKey, pubKey, err := GeneratePubPriKeyPair() 
	if err != nil {
		return []byte(""), []byte(""), err
	}

	privateKey := x509.MarshalPKCS1PrivateKey(priKey)
	publicKey, err := x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		return nil, nil, err
	}
	return privateKey, publicKey, nil
}