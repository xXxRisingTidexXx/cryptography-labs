package cryptolabs_test

import (
	"github.com/xXxRisingTidexXx/cryptography-labs"
	"testing"
)

func TestTrithemius2(t *testing.T) {
	testCipher(
		t,
		cryptolabs.NewTrithemius2(cryptolabs.DefaultAlphabet),
		spec{"", "", ""},
		spec{
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut " +
				"labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco" +
				" laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit " +
				"in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat c" +
				"upidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			"",
			"Lpthq?owАГwDpБzГЄJИБЙNwЗВОUTҐННУИЗЧЙЩЮЩfКОФЯЦґТЩвЩqЩґвкxwмггАднГзкчцсфкІяочюю,Ош.фюч щGAHЩKKЯ" +
				"E,!KOCеEUзGSQUYMмWLSaOтQcajoVььSs fplq?gkAvsyuzGИuҐАtЕPOЙОЄОTЛНТФУЧІbКбМЬМУвНґЧббoивгЦеЬи" +
				"wзЯблпітҐплцнЗююЇкцф ?ч,РцFУщцЦьCBCF-Hб?KKQDQVCWїїфbQbнLggSтYimkYшZljnr.kq!wkwznrpzqsЕyИK" +
				"БЄNНІЗПЛРВТЖXЧІПНЩdМЮЯПiПЦЬЮжбpЧеґжїаwвпzетзїдхЕсщстїЇщл хо?\"?ТТвHшюCH.KIЯLCIPґMBCBGFEYй" +
				"JcYSOMgOiуeggчknlgceov?!xАuБBsyEpЖАДrKЗКГOІГҐЖГИГWЖИХЇЦЬФЮfЦЩЧШЦдmСвЮгrбЬuалнyібгнсхоЄ",
		},
		spec{
			"AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
			"",
			"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyzАБВГҐДЕЄЖЗИІЇЙКЛМНОПРС",
		},
	)
	testCipher(
		t,
		cryptolabs.NewTrithemius2(cryptolabs.UkrainianAlphabet),
		spec{"ПАРОЛЬ", "", "ПБТСПВ"},
		spec{
			"ТИХАВОДАГРЕБЛЮРВЕГОЛЦІНІДЕВПАСТИТИХІШЕВОДИНИЖЧЕТРАВИЯБЛУКОВІДЯБЛУНІ",
			"The key",
			"ТІЧГЕУІЄІЩМЇЧІБНТРГБЙЯЕБЩЮШЙШНПЖСИЦЙЮИЄФЇОЦРПЕПГВМПШОСГЇҐЗЧҐЯЧЬЇСМІ",
		},
	)
}
