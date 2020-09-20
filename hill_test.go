package cryptolabs_test

import (
	"github.com/xXxRisingTidexXx/cryptography-labs"
	"testing"
)

func TestHill(t *testing.T) {
	testCipher(
		t,
		cryptolabs.NewHill(cryptolabs.DefaultAlphabet),
		&spec{"", "", ""},
		&spec{"Abeceda", "", "Abeceda"},
		&spec{"", "нАдіЙний клЮч", ""},
		&spec{"Невалідний клю", "-sjfdls", "Невалідний клю"},
		&spec{
			"The protagonist of Hamlet is Prince Hamlet of Denmark, son of the recently deceased King Haml" +
				"et, and nephew of King Claudius, his father's brother and successor. Claudius hastily mar" +
				"ried King Hamlet's widow, Gertrude, Hamlet's mother, and took the throne for himself.fff",
			"BB!BBBABB",
			"ЕЬІшФЯтцПqцЧY?ґiІПVZhЩрЛbПaнzHvюСдyWHПЇNчУiІПBZhкоЇоСeCФЯ?ПgиМlиГWвмЕPщЯ.ЧqСвЕGлОеzVOЮС!f\"qм" +
				"УкНnщДЙяЖfЙпТхОoоІXOЮСQa,FфРFфШОeчRцЦtxДЯуІяаqкИОЧ ЦрЙjУзМпЬдИбЕВCЯМcцЗЙІЙфІ!ЬkJоОДюПQШeX" +
				"чбТжЄ ksоІYHПЇNчУТikєцІчЬqEckqFгЩрЄЯ.\"qмУЮСrKТeЧ ЦїЛlщДЙжКlИюФfОЦзЛlqюЮмЗWZщЮPЄЛгчТіЗYЇж" +
				"И",
		},
	)
	testCipher(
		t,
		cryptolabs.NewHill(cryptolabs.EnglishAlphabet),
		&spec{"ACT", "GYBNQKURP", "POH"},
		&spec{"YIUQ", "ABDB", "ICQY"},
		&spec{"DETERMINANTS", "DCIH", "RANYXMYZANPS"},
	)
}
