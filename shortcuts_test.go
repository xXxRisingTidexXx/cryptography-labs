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

func testSigner(t *testing.T, signer cryptolabs.Signer, docs ...doc) {
	for i := range docs {
		if !signer.Verify(signer.Sign(docs[i].m, docs[i].d), docs[i].e) {
			t.Errorf("cryptolabs_test: got invalid signature for %v", docs)
		}
	}
}
