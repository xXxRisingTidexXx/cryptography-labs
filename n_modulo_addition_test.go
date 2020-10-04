package cryptolabs_test

import (
	"github.com/xXxRisingTidexXx/cryptography-labs"
	"testing"
)

func TestNModuloAddition(t *testing.T) {
	testCipher(
		t,
		cryptolabs.NewNModuloAddition(cryptolabs.DefaultAlphabet),
		spec{"GRONSFELD", "Appveyorxnhjx892d1", "GЕГЖwГsВА"},
		spec{
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut " +
				"labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco" +
				" laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit " +
				"in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat c" +
				"upidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			"xnZjdqjd,.pska.S'---.EMSQ.WDSMLZZL",
			"ЖЧМЙЛiНОmnЧkЙКeЕpшndm!mҐumQ-uАyНГnзНОЩРiЗЕciУгЇЖgyюZgdm'EИuWOgЕEpЄПГвЧВbТРСОikhЦУВbvgYpim!ЕІI" +
				"eweЕГpRГҐpМЙРНаЙVgTСЮИSTГglpVч!gІIXЗlҐElВRxЮЦЄСVґЙМcTЧmcМnАqшijlxГЇtцААwГnЄОlзСЙТVгРКUfНЯ" +
				"cІTtmmdnцruИyцМwKmwЄЛДЮШRЙЦiЙГчVЩЬТКWЕюXjiliВЇqmP-VЕtНRlиЯГbІаЩРYцОЯСКkKgiшmXtГwxXЗgwГtОR" +
				"tгfРУКгФТUmПkЮҐeАrшZnliEuyeЄxҐEoЙИzєНRЙУiКУabЛґcЙnГjVшkTvusЗnЙ!KQЖБГАзНПЦVбНМnцЩРЇАXuYoшX" +
				"ntuvqmwwKzzЇRАєЧЄЇЄЩШXчlвЮЩSbДюXpgieEЖИbOrxrtБЄlpМГЧЄаЩМnцЧЯСІbІюVidf!uvIXКwKxlАЙВиХS",
		},
		spec{
			"aaaAAAAAaaaaaAAAaaaaaAAAAaaaAAAaaaAAAAAAAAAAaaaAAaaaaAAAAAaaAAaaAAAAaAAa",
			"djndowsjnakjsdnjasbdibdhiashdioasd--jdsskaassaaBHJKKkaoisdj",
			"ГЗЙdowsjЙАИЗОdnjАОБГЖbdhiАОЄdioАОГ--jdsskaasОААBHjkkИaoisdЗГjnГКwsjnАkjО",
		},
		spec{
			"Думи мої, думи мої, Лихо мені з вами! Нащо стали на папері Сумними рядами?.. Чом вас вітер не" +
				" розвіяв В степу, як пилину? Чом вас лихо не приспало, Як свою дитину?...",
			"Результати Оприлюднено",
			"DСНХєдЯҐмвЬАЧТвУжЗіЮrРҐМбЮЙеХХУвЧdЧТеєЄxдМзЮБyМХєеКХаВіjЮУвvйЙЦЖХРЇОаРЄдФЬлгхXЩШвЗУОзАРЯlОбЯЙ" +
				"тбИУҐгЗЙїXєіПЛНЯїЇЬЛлЦЯЩВЮХьИДЬОєХxЬЮФРҐМбЯЙтаЇФУіdЦЬґєТІзПІШИЮЕХЩЯЮЛпгцІ",
		},
	)
}
