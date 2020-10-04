package cryptolabs_test

import (
	cryptolabs "github.com/xXxRisingTidexXx/cryptography-labs"
	"testing"
)

func TestFeistel(t *testing.T) {
	testCipher(
		t,
		cryptolabs.NewFeistel(
			func(_, text, key rune) rune {
				return text + key
			},
		),
		spec{"", "", ""},
		spec{"Hello, World!", "", "Hello, World!"},
		spec{"", "Привіт, світ!", ""},
		spec{"", "Привіт, світ!", ""},
		spec{"hello", "qwerty", "{U\x1dZŸ"},
	)
}
