package cryptolabs_test

import (
	"github.com/xXxRisingTidexXx/cryptography-labs"
	"testing"
)

func TestGOST341094(t *testing.T) {
	signer, m := cryptolabs.NewGOST341094(67, 11, 25, 6, 9), int64(31)
	if s := signer.Encode(m, 0); !signer.Verify(s, m) {
		t.Errorf("cryptolabs_test: got unverification for s = %d", s)
	}
	if !cryptolabs.NewExplicitGOST341094(67, 11, 25, 24, 3).Verify(5, 1) {
		t.Errorf("cryptolabs_test: got unverification")
	}
	if !cryptolabs.NewExplicitGOST341094(67, 11, 25, 24, 4).Verify(3, 1) {
		t.Errorf("cryptolabs_test: got unverification")
	}
	if cryptolabs.NewExplicitGOST341094(67, 11, 25, 24, 4).Verify(5, 1) {
		t.Errorf("cryptolabs_test: got verification")
	}
}
