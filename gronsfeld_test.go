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
			"Nosmp!pwxВq-dqmxz\"vkt.ipiАB?ksqsgdВmАxt blltpzhqrj gmrБB-ueeAgs\"lnВwpof.Вmtsqr.qqgpknlyqt,v" +
				"ВAsddosm-iА\"iwprrg.vinqc btluБh'AYw goru\"df nqqmt\"Аmrlao-ByБlu owvxyБiAiАetdrБhwkooAxp" +
				"shrks-lccxzpv,njАl!БА?iplqwjyAlА,ebAfstttls-cqoБmxxct,AGypz?iywe,jАВyh,dptrv\"psAvhptfqmu" +
				"ggrjБ-mu\"АwpxpvbВm\"ygljБ-izzjAgllnvvAkrnosm-iБ\"kВklav.wВsoc qiumhАzz?-EzdnxАhwr.АlrА\"" +
				"tkgdeebВAjxrieiweА\"swr-ptprllqv,.АxrА\"nv!funqjAxxk pnimjpfAhhsgsГvА-oomtlx\"hsqq-if.nАА" +
				"-nacwuytA",
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
			"Еяои-хцм--йянй.хрї?BУкцсAхєрїBі дзфк?-Хзьс.ьфаопAсб-чзрзср,СцххкнйAщ жбхї?!CAЮппAібф.ійтзщAсє" +
				"-шчидїAґ ДBщцєтюD.,лBсиопхч'-вчн-гзу оп т.ркBруіьсаочC!анAьгсяBєихпхч'!BC",
		},
	)
}
