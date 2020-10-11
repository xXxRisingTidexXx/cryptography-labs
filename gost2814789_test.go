package cryptolabs_test

import (
	"github.com/xXxRisingTidexXx/cryptography-labs"
	"testing"
)

func TestNewGOST2814789(t *testing.T) {
	testCipher(
		t,
		cryptolabs.NewGOST2814789(cryptolabs.CBRF),
	)
}
