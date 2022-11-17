package signing

import (
	"crypto"
	"crypto/rsa"
	"crypto/rand"
	"crypto/sha512"
)



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
