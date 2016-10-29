package decryptor

import (
	"encoding/binary"
	"errors"
	"io/ioutil"
)

const KeySize = 32

func Decrypt(text, key []byte) ([]byte, error) {
	if len(key) > KeySize {
		return nil, errors.New("len(key) must be <= 32")
	}

	var k [KeySize]byte
	copy(k[:], key)

	n := len(text)
	t := make([]byte, n, n+(4-n%4))
	copy(t, text)

	f := func(x uint32) uint32 {
		return x*134775813 + 1
	}
	res := make([]byte, n, n+(4-n%4))
	for i, p := 0, uint32(0); i < n; i += 4 {
		x := binary.LittleEndian.Uint32(key[i%KeySize : (i%KeySize)+4])
		p = f(x ^ p)

		x = binary.LittleEndian.Uint32(t[i : i+4])
		binary.LittleEndian.PutUint32(res[i:i+4], p^x)
	}

	return res[:n], nil
}

func DecryptFiles(text, key string) ([]byte, error) {
	k, err := ioutil.ReadFile(key)
	if err != nil {
		return nil, err
	}

	t, err := ioutil.ReadFile(text)
	if err != nil {
		return nil, err
	}

	return Decrypt(t, k)
}
