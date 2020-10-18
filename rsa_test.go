package cryptolabs_test

import (
	"github.com/xXxRisingTidexXx/cryptography-labs"
	"testing"
)

func TestRSA(t *testing.T) {
	testCipher(
		t,
		cryptolabs.NewRSA(5, 13, 37),
	)
}
