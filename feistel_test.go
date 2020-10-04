package cryptolabs_test

import (
	cryptolabs "github.com/xXxRisingTidexXx/cryptography-labs"
	"testing"
)

func TestFeistel(t *testing.T) {
	testCipher(
		t,
		cryptolabs.NewFeistel(),
	)
}
