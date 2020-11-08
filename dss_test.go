package cryptolabs_test

import (
	"github.com/xXxRisingTidexXx/cryptography-labs"
	"testing"
)

func TestDSS(t *testing.T) {
	signer, m := cryptolabs.NewDSS(251, 25, 17, 11, 113), int64(17)
	if s := signer.Encode(m, 0); !signer.Verify(s, m) {
		t.Errorf("cryptolabs_test: got unverification for s = %d", s)
	}
	signer, m = cryptolabs.NewDSS(251, 125, 20, 7, 21), int64(51)
	if s := signer.Encode(m, 0); !signer.Verify(s, m) {
		t.Errorf("cryptolabs_test: got unverification for s = %d", s)
	}
	signer = cryptolabs.NewExplicitDSS(271, 27, 106, 28, 7)
	if !signer.Verify(3, 17) {
		t.Errorf("cryptolabs_test: got unverification")
	}
	signer = cryptolabs.NewExplicitDSS(271, 27, 106, 28, 13)
	if signer.Verify(8, 17) {
		t.Errorf("cryptolabs_test: got verification")
	}
	signer = cryptolabs.NewExplicitDSS(271, 27, 106, 28, 17)
	if !signer.Verify(10, 17) {
		t.Errorf("cryptolabs_test: got unverification")
	}
}
