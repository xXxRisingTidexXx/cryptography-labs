package cryptolabs_test

import (
	cryptolabs "github.com/xXxRisingTidexXx/cryptography-labs"
	"testing"
)

func TestFeistel(t *testing.T) {
	testCipher(
		t,
		cryptolabs.NewFeistel(cryptolabs.Magma),
		spec{"", "", ""},
		spec{"Hello, World!", "", "Hello, World!"},
		spec{"", "Привіт, світ!", ""},
		spec{"", "Привіт, світ!", ""},
		spec{"helloh", "qwerty", "{U\x1dZŸ\u008e"},
		spec{"hello", "qwerty", "hello"},
		spec{
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor i",
			"*&Yhdhqwid23J((@hdksak__K9239",
			"ԍڇϯȣ\x14ƹƬɝ˛Ɣ\x14ƹō?ˍʧ΄ȔƏ̨̄ǚ\aÚĈ̤͝ƲA\x06ƵƉʑƈ[ĵɋۅ΄Ṵ̏\u05c9ƬɝƁĬďŸƑ½\u0083Iķˑ̠ўÙƏȸɍŀǲƜW٫ϒƁ̈΅ɬ҄ϊ[ĵǄ)ª\x11دϕ",
		},
		spec{
			"Думи мої, думи мої, Лихо мені з вами! Нащо стали на папері Сумними рядами?.. Чом",
			"І день йде, і ніч минає.",
			"ᴓ▒ἁ⚏ဍᓋ᰽◦㋡⬇ᰳ⒒ἁ⚏ဍᓋ᰽◦㋡⬇ἠ◵ᶎ\u245fဍᓋ᳴┖\u1C93⧃\u175cᔅᶟ\u2433ἁ⚏㋞⪕ᰚ\u243dὢ⚟ፋᘔᶯ✳ᵀ⛵ᆛᓘᰵ⪏᱔⑵ἶ♚ᴅ▹ኋ" +
				"឴ᵘ⚡࿚Ǖἁ⚏ፅᗯᵔ♁ᴝ⑫\u1ccf⥖\u31eb⤽့ᓞ\u0fe3Ǘ",
		},
		spec{
			"Я на гору круту кам'яную буду камінь важкий підіймать. І несучи вагу ту страшную буду пісню.",
			"Людина, нібито, не літає... А крила має!",
			"\u1aeeᚿՌݝˌࠐսІᣚᐫՑي\u0530ݻᣚᐫ١ݪ᭜ᗬԓЇД՛˾ର݆ڛᣚᐫ١ݪԥӔԸӍ\x05ࠠعנݩ\u05ca᧰ᗑۨۿ\u0a5dঌֱ֔хАڭƴȹଐŢ\u0880եօםգᥟᔸѹ" +
				"ҢԉדƕॠᣚᐫТہקڐ\u09baʈД՛˾ର݆ڛᣚᐫۨۿǍ\u0a61ᡣᑖ",
		},
	)
}
