package cryptolabs_test

import (
	"github.com/xXxRisingTidexXx/cryptography-labs"
	"testing"
)

func TestMagicSquare(t *testing.T) {
	testCipher(
		t,
		cryptolabs.NewMagicSquare(),
		//&spec{"", "", ""},
		//&spec{"ВІРТУАЛЬНИЙ КАНАЛ", "16,3,2,13,5,10,11,8,9,6,7,12,4,15,14,1", "SHIT"},
	)
}
