package cryptolabs_test

import (
	"github.com/xXxRisingTidexXx/cryptography-labs"
	"testing"
)

func TestGronsfeld(t *testing.T) {
	testCipher(
		t,
		cryptolabs.NewGronsfeld(cryptolabs.DefaultAlphabet),
		&spec{"", "", ""},
		&spec{"Abeceda", "", "Abeceda"},
		&spec{"", "нАдіЙний клЮч", ""},
		&spec{"Невалідний ключ", "-sjfdls", "Невалідний ключ"},
		&spec{"GRONSFELD", "2015", "IRPSUFFQF"},
		&spec{
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut " +
				"labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco" +
				" laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit " +
				"in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat c" +
				"upidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			"2018347758430219873",
			"Nosmp!pwxВq-dqmxzBvkt.ipiАD?ksqsgdВmАxt blltpzhqrj gmrБD-ueeCgsBlnВwpof.Вmtsqr.qqgpknlyqt,vВC" +
				"sddosm-iАBiwprrg.vinqc btluБhACYw goruBdf nqqmtBАmrlao-DyБlu owvxyБiCiАetdrБhwkooCxpshrks" +
				"-lccxzpv,njАl!БА?iplqwjyClА,ebCfstttls-cqoБmxxct,CGypz?iywe,jАВyh,dptrvBpsCvhptfqmuggrjБ-" +
				"muBАwpxpvbВmBygljБ-izzjCgllnvvCkrnosm-iБBkВklav.wВsoc qiumhАzz?-EzdnxАhwr.АlrАBtkgdeebВCj" +
				"xrieiweАBswr-ptprllqv,.АxrАBnv!funqjCxxk pnimjpfChhsgsГvА-oomtlxBhsqq-if.nАА-nacwuytC",
		},
		&spec{
			"aaaAAAAAaaaaaAAAaaaaaAAAAaaaAAAaaaAAAAAAAAAAaaaAAaaaaAAAAAaaAAaaAAAAaAAa",
			"0123456789",
			"abcDEFGHijabcDEFghijaBCDEfghIJAbcdEFGHIJABCDefgHIjabcDEFGHijABcdEFGHiJAb",
		},
		&spec{
			"Думи мої, думи мої, Лихо мені з вами! Нащо стали на папері Сумними рядами?.. Чом вас вітер не" +
				" розвіяв В степу, як пилину? Чом вас лихо не приспало, Як свою дитину?...",
			"19203984138913",
			"Еяои-хцм--йянй.хрї?DУкцсCхєрїDі дзфк?-Хзьс.ьфаопCсб-чзрзср,СцххкнйCщ жбхї?!ECЮппCібф.ійтзщCсє" +
				"-шчидїCґ ДDщцєтюF.,лDсиопхчA-вчн-гзу оп т.ркDруіьсаочE!анCьгсяDєихпхчA!DE",
		},
	)
}
