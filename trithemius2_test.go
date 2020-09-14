package cryptolabs_test

import (
	"github.com/xXxRisingTidexXx/cryptography-labs"
	"testing"
)

func TestTrithemius2(t *testing.T) {
	testCipher(
		t,
		cryptolabs.NewTrithemius2(cryptolabs.DefaultAlphabet),
		&spec{"", "", ""},
		&spec{"Abeceda", "", "Abeceda"},
		&spec{"", "", ""},
		&spec{"Невалідний ключ", "", "Невалідний ключ"},
		&spec{"trithemius2", "2020", ""},
	)
	testCipher(
		t,
		cryptolabs.NewTrithemius2(cryptolabs.UkrainianAlphabet),
		&spec{"ПАРОЛЬ", "", "ПБТСПВ"},
	)
}
