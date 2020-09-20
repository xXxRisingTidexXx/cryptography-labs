package cryptolabs_test

import (
	"github.com/xXxRisingTidexXx/cryptography-labs"
	"testing"
)

func TestNModuloAddition(t *testing.T) {
	testCipher(
		t,
		cryptolabs.NewNModuloAddition(cryptolabs.DefaultAlphabet),
	)
}
