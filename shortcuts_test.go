package cryptolabs_test

import (
	"github.com/xXxRisingTidexXx/cryptography-labs"
	"testing"
)

func testCipher(t *testing.T, cipher cryptolabs.Cipher, specs ...spec) {
	for i := range specs {
		if output := cipher.Encrypt(specs[i].text, specs[i].key); output != specs[i].output {
			t.Errorf("cryptolabs_test: got invalid encrypted texts, %s != %s", output, specs[i].output)
		}
		if text := cipher.Decrypt(specs[i].output, specs[i].key); text != specs[i].text {
			t.Errorf("cryptolabs_test: got invalid decrypted texts, %s != %s", text, specs[i].text)
		}
	}
}

func testEncode(t *testing.T, signer cryptolabs.Signer, docs ...doc) {
	for i := range docs {
		s := signer.Encode(docs[i].m, docs[i].d)
		if signer.Decode(s, docs[i].e) != docs[i].m {
			t.Errorf("cryptolabs_test: got invalid encoding for %v, %d", docs[i], s)
		}
	}
}

func testDecode(t *testing.T, signer cryptolabs.Signer, signs ...sign) {
	for i := range signs {
		if m := signer.Decode(signs[i].s, signs[i].e); m != signs[i].m {
			t.Errorf("cryptolabs_test: got invalid decoding for %v, %d", signs[i], m)
		}
	}
}
