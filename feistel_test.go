package cryptolabs_test

import (
	cryptolabs "github.com/xXxRisingTidexXx/cryptography-labs"
	"testing"
)

func TestFeistel(t *testing.T) {
	testCipher(
		t,
		cryptolabs.NewFeistel(
			func(text, key rune) rune {
				return text + key
			},
		),
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
			"В дверях эдема ангел нежный Главой поникшею сиял, А демон мрачный и мятежный Над адской бездн" +
				"ою летал. Дух отрицанья, дух сомненья На духа чистого взирал И жар невольный умиленья Впе",
			"Людина, нібито, не літає... А крила має!",
			"ᣉᓢਡবմЕЫԷᵒᵀЖ҄ቛ㞔\u0087\u0870૧अշֽŢ\u0880ւݍЏՕ᧰ᗑБ֓ՕՐլԞ`ࢰីพߕՈؒר\u1aedᗖրۉࢽ૿഻ॄ᪇ᝀЖ્҄ऄ\u193c\u169d֫֔҈րЏՕ᧰ᗑ" +
				"ᥟᔸݴ۬՜ٲ׀֖ᥳ⪣\x02ࡀ֫Ґ\u0087\u0870֢ܴᅏ㟊᧰ᗑҥё\u0aba߷ڂڕ\u1aedᗖӏۣЩڲ\u19dcᔳɛऀሇ᧫١ΐәّْ߈Ғ\u0590քԼ഻ॄڠܤᣄᓅ۶ֱ" +
				"\u05f6ٌډ֕քԼ\x02ࡀᣧᕠڠܤѴӅ̸ৰۆרѧڒ״ҋ\x05ࠠՎԧקڐᔲ⇣\u193fᗘץ؞᧷ᐰ\u0aa9्ӷЂֆ۳ЏՕ᧰ᗑञફ٬֨ډ֕քԼ¥ठ⥫\u1f47",
		},
	)
}
