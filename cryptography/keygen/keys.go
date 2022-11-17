package keygen

import (
	"crypto/rsa"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
)


// Generate keys that can be used immediately.
func GeneratePubPriKeyPair() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	priKey, err := rsa.GenerateKey(rand.Reader, 4096) 
	if err != nil {
		return nil, nil, err
	}
	// The public key is a part of the *rsa.PrivateKey struct
	// so just extract public key component
	pubKey := &priKey.PublicKey

	return priKey, pubKey, nil
}


// Generate private/public keys in bytes
func GeneratePubPriKeyPairInPem(keyFile string) error {
	priKey, pubKey, err := GeneratePubPriKeyPair() 
	if err != nil {
		return err
	}

	// Encode private key to PKCS#1 ASN.1 PEM.
	priKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type: "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(priKey),
		},
	)

	// Encode public key to PKCS#1 ASN.1 PEM.
	pubKeyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type: "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(pubKey),
		},
	)

	// Write private key to file
	err = ioutil.WriteFile(keyFile+".rsa", priKeyPEM, 0700)
	if err != nil {
		return err
	}

	// Write public key to file
	if err = ioutil.WriteFile(keyFile+".rsa.pub", pubKeyPEM, 0755); err != nil {
		return err
	}
	return nil
}