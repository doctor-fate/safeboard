package decryptor

import (
	"testing"
)

func TestDecryptLongKey(t *testing.T) {
	text := []byte("The power to protect what matters most")
	key := make([]byte, 33)

	_, err := Decrypt(text, key)
	if err == nil {
		t.Fatal("expected \"len(key) > 32\" error, but got nil instead")
	}
}

func TestDecrypt(t *testing.T) {
	text := []byte("The power to protect what matters most")
	key := []byte("Kaspersky Lab")
	expected := []byte{
		0x2C, 0xFA, 0xA6, 0x32, 0xE2, 0x3B, 0x9C, 0x19, 0xEA, 0x54, 0x92, 0x84, 0xC3,
		0x40, 0xD3, 0x88, 0x04, 0x65, 0x10, 0x36, 0x11, 0xB5, 0x90, 0x7A, 0x82, 0x2E,
		0xE8, 0x08, 0xBB, 0x56, 0x65, 0xA1, 0xE6, 0x41, 0x5D, 0xB3, 0xC2, 0x57,
	}

	res, err := Decrypt(text, key)
	if err != nil {
		t.Fatal(err)
	}

	if len(res) != len(expected) {
		t.Errorf("len of expected sequence is %d, but len of result is %d instead", len(res), len(expected))
	}

	for i := 0; i < len(res); i++ {
		if res[i] != expected[i] {
			t.Fatalf("expected sequence is\n%X,\nbut res is\n%X\ninstead", expected, res)
		}
	}
}
