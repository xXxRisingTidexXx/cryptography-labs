package cryptolabs_test

import (
	"github.com/xXxRisingTidexXx/cryptography-labs"
	"testing"
)

func TestHill(t *testing.T) {
	testCipher(
		t,
		cryptolabs.Hill,
		&spec{"", "", ""},
		&spec{"Abeceda", "", "Abeceda"},
		&spec{"", "нАдіЙний клЮч", ""},
		&spec{"Невалідний ключ", "-sjfdls", "Невалідний ключ"},
		&spec{"HILL", "2020", ""},
	)
}
