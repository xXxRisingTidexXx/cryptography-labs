package cryptolabs_test

import (
	"github.com/xXxRisingTidexXx/cryptography-labs"
	"testing"
)

func TestElgamal(t *testing.T) {
	testCipher(
		t,
		cryptolabs.NewElgamal(31, 14, 17, 16),
		//spec{"", "", ""},
		//spec{"Петраківський", "", "\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11\x11" +
		//	"\x11\x11\n\x15\x10\x0f\a\f\x03\b\x04\x15\f\v\x00"},
		//spec{"Belvedeur", "", "1\n\u0015\u0010\u000F\a\f"},
		//spec{"Hello, world!, Shut up and just let me calm down", "", "1" +
		//	"\n\u0015\u0010\u000F\a\fZwx\u000F"},
	)
}
