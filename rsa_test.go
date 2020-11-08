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
	testEncode(t, cryptolabs.NewSignerRSA(91), doc{7, 29, 5})
	testEncode(t, cryptolabs.NewSignerRSA(55), doc{13, 27, 3})
	signer := cryptolabs.NewSignerRSA(93)
	testDecode(t, signer, sign{12, 11, 21}, sign{67, 11, 25})
	if signer.Decode(29, 11) == 9 {
		t.Errorf("cryptolabs_test: got false positive decoding")
	}
}
