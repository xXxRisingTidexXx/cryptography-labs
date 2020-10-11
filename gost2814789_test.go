package cryptolabs_test

import (
	"github.com/xXxRisingTidexXx/cryptography-labs"
	"testing"
)

func TestNewGOST2814789(t *testing.T) {
	testCipher(
		t,
		cryptolabs.NewGOST2814789(
			cryptolabs.NewSBox(
				[8][16]byte{
					{4, 10, 9, 2, 13, 8, 0, 14, 6, 11, 1, 12, 7, 15, 5, 3},
					{14, 11, 4, 12, 6, 13, 15, 10, 2, 3, 8, 1, 0, 7, 5, 9},
					{5, 8, 1, 13, 10, 3, 4, 2, 14, 15, 12, 7, 6, 0, 9, 11},
					{7, 13, 10, 1, 0, 8, 9, 15, 14, 4, 6, 12, 11, 2, 5, 3},
					{6, 12, 7, 1, 5, 15, 13, 8, 4, 10, 9, 14, 0, 3, 11, 2},
					{4, 11, 10, 0, 7, 2, 1, 13, 3, 6, 8, 5, 9, 12, 15, 14},
					{13, 11, 4, 1, 3, 15, 5, 9, 0, 10, 14, 7, 6, 8, 2, 12},
					{1, 15, 13, 0, 5, 7, 10, 4, 9, 2, 3, 14, 6, 11, 8, 12},
				},
			),
		),
		spec{},
		spec{"", "hello, world,:)!", ""},
		spec{"hello, world,:)!", "", "hello, world,:)!"},
		spec{"hello, world,:)!", "2123", "hello, world,:)!"},
		spec{
			"hello, world,:)!",
			"3d91__23easHJsjbc-+j3ddnscjhs-32",
			"\x1e!\xfa\xea\x93\xcf!\xd8\xe35\x19$\x00ҝ;",
		},
		spec{
			"Криптографія - річ, що на найбільш явному місці тримає те, що має бути якнайглибше " +
				"приховано.",
			"9u9&87w3hdskJHGjhb_i1ocsnkjacj0<",
			"jy\x17\x83\xde{O`\xd2e\xbf\xed\x1e\xb0\xea\xcc\x03\\&\xea3\xa5\x9aW\x90\xe45h\x17" +
				"\x10@\x87I\x92\x8fb\x8aB\xf3{a?ωH}s\x13\xc7RV\x8e\xfbo\xf2\x14\x9c\xd0L\xc6a" +
				"\x12r\xa1\bwԟ\xfd\x88Q.\"_\xfd\xef\x13\xf4\x1e$a%\xcc\xef\a\xdcԹ\xb6u{Ы\xb6" +
				"\xf3Ȭ\x82Xt\x8b\x03\nY\x8e\xa2E\xfa\xbe\x02<_|\x9b \xa6\x05\xfb`P\x81\n->\xd7" +
				"\xcei\xf3O\xae%&\xcab\xd8\x1d\xe8\xe7T-8\x8dI:ɲD\x99\xd9N\x95\x86%k\x8f\xef\xfd" +
				"\x96tN\xac\xc5\xd4\x10{\f\x1a\xc8",
		},
	)
}
