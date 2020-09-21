package cryptolabs_test

import (
	"github.com/xXxRisingTidexXx/cryptography-labs"
	"testing"
)

func TestXOR(t *testing.T) {
	testCipher(
		t,
		cryptolabs.NewXOR(),
		&spec{"", "", ""},
		&spec{
			"ЗАГРОЗА ІСНУЄ ЗАВЖДИ І ВСЮДИ",
			"КЛЮЧІ ВІД ТВОГО СЕРЦЯ ВЖЕ ДАВНО В МОЇХ РУКАХ",
			"\r\v=\a\x18з\x02Ц\x12Ё?1\x1aг\tа3\x034>ЏЦв\x044Ў\x00\b",
		},
		&spec{
			"ЗАГРОЗА ІСНУЄ ЗАВЖДИ І ВСЮДИ",
			"КОРОТЕСЕНЬКИЙ КЛЮЧ",
			"\r\x0e3><\x021е\x1b\r\a;\x1d\x00\r\v<1\x0e\x06Ѐ\x18Ђ\a\x00;\t4",
		},
		&spec{
			"За горами гори, хмарою повиті, Засіяні горем, кровію политі.",
			"Nniu29ncsdkjnc i2893e)(U)23jodjT!YT!asjhbdnkjJ*hckhsdnk*&(Hsdacnkqw",
			"љўIцЌѹўџыDјєЮћ\fIѷЄЉѳћѧ\bѪЗЀЋШйHJуБИЂѮќХJћќФћїFjАШѝљонDёѕБОѪО]",
		},
		&spec{
			"The protagonist of Hamlet is Prince Hamlet of Denmark, son of.",
			"??>mLmkoidjf(*(8329djoH!YHuhwedhc  jkwsdal!L~I`wpqwed23dxF34ggfhk",
			"kW[M<\x1f\x04\x1b\b\x03\x05\bAY\\\x18\\T\x19,\v\x02$D-h\x1c\x1bW5\x16\x01\rCEJ#\x16\x1e\b\x04" +
				"\x18\x01#\x18i$\x12\x1e\x1c\x16\x17\x0f\x1e\x13\x17\x17(\x13[\x01I",
		},
		&spec{
			"Вічний революціонер — Дух, що тіло рве до бою, Рве за поступ, щастя й волю.",
			"  KNOoqw0)(*@8328eu334%THTrhglc/z/xc.zx kz;dcpweo*(*jsaowdewfew)))(Ud3dkqjsakxb",
			"вѶЌѳѷіQзЅЛЖБЎѾѥЌЅѐе\x13‧\x14бЗЍxRСљLСѹсБXУМяXДѕZЊњЭ\\WхѝП\bНњSўёжЦЦшJEоЙѨѫѧuѝ\x13іѕъФ]",
		},
		&spec{
			"COVID-19 вже надоїв. Хочемо в університет.",
			"NIKHNIUH87jh9ih nv_))O09jf3rihjefsdfhnk",
			"\r\x06\x1d\x01\nddq\x18Ѕќѝ\x19єјДѐСѭ\a\tѪЎѾџњЍRћHЩјасёЦЩіЩѻЋe",
		},
	)
}
