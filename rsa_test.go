package cryptolabs_test

import (
	"github.com/xXxRisingTidexXx/cryptography-labs"
	"testing"
)

func TestRSA(t *testing.T) {
	testCipher(
		t,
		cryptolabs.NewRSA(4651, 7919, 6113),
		spec{"", "", ""},
		spec{"TE", "", "NRVBMB"},
		spec{"CQ", "", "UDKZXB"},
		spec{"HELLO", "", "OVDRG"},
		spec{"DAYLO", "", "YXCFY"},
	)
}
