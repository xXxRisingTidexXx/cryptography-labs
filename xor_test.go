package cryptolabs_test

import (
	"github.com/xXxRisingTidexXx/cryptography-labs"
	"testing"
)

func TestXOR(t *testing.T) {
	testCipher(t, cryptolabs.NewXOR())
}
