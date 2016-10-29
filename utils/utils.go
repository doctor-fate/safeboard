package utils

import (
	"io/ioutil"

	"github.com/doctor-fate/safeboard/decryptor"
)

func WriteKeyAndDecrypt(key []byte, kpath, epath, dpath string) error {
	if err := ioutil.WriteFile(kpath, key, 0644); err != nil {
		return err
	}

	dec, err := decryptor.DecryptFiles(epath, kpath)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(dpath, dec, 0644); err != nil {
		return err
	}

	return nil
}
