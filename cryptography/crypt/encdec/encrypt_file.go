package encdec

import (
	"io/ioutil"
	"errors"
	"fmt"
)


func EncryptFile(srcFpath string, dstFpath string, keyFpath string) error {
	plainTxt, err := ioutil.ReadFile(srcFpath)
	if err != nil {
		errMsg := fmt.Sprintf("File read err: %v", err.Error())
		return errors.New(errMsg)
	}

	// Read the secret key which feeds into AES cihper block
	key, err := ioutil.ReadFile(keyFpath)
	if err != nil {
		errMsg := fmt.Sprintf("File read err: %v", err.Error())
		return errors.New(errMsg)
	}

	cipherTxt := EncryptTxt(plainTxt, string(key))

	// save the cipher text 
	if err = ioutil.WriteFile(dstFpath, cipherTxt, 0777); err != nil {
		errMsg := fmt.Sprintf("File write err: %v", err.Error())
		return errors.New(errMsg)
	}

	return nil
}


func DecryptFile(srcFpath string, dstFpath string, keyFpath string) error {
	cipherTxt, err := ioutil.ReadFile(srcFpath)
	if err != nil {
		return err
	}

	key, err := ioutil.ReadFile(keyFpath)
	if err != nil {
		return err
	}

	plainTxt := DecryptCipherTxt(cipherTxt, string(key))

	if dstFpath == "" {
		fmt.Println(string(plainTxt))
	} else {
		if err := ioutil.WriteFile(dstFpath, plainTxt, 0777); err != nil {
			return err
		}
	}
	
	return nil

}