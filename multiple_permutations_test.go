package cryptolabs_test

import (
	"github.com/xXxRisingTidexXx/cryptography-labs"
	"testing"
)

func TestMultiplePermutations(t *testing.T) {
	testCipher(
		t,
		cryptolabs.NewMultiplePermutations(),
		&spec{"", "", ""},
		&spec{
			"ЗАГРОЗА ІСНУЄ ЗАВЖДИ І ВСЮДИ",
			"",
			"",
		},
		&spec{
			"За горами гори, хмарою повиті, Засіяні горем, кровію политі. Споконвіку Прометея Там орел кар" +
				"ає, Що день божий добрі ребра Й серце розбиває. Розбиває, та не вип’є Живущої крові — Вон" +
				"о знову оживає І сміється знову. Не вмирає душа наша, Не вмирає воля. І неситий не виоре " +
				"На дні моря поле. Не скує душ",
			"",
			"",
		},
		&spec{
			"The protagonist of Hamlet is Prince Hamlet of Denmark, son of the recently deceased King Haml" +
				"et, and nephew of King Claudius, his father's brother and successor. Claudius hastily mar" +
				"ried King Hamlet's widow, Gertrude, Hamlet's mother, and took the xx",
			"",
			"",
		},
		&spec{
			"Вічний революціонер — Дух, що тіло рве до бою, Рве за поступ, щастя й волю, — Він живе, він щ" +
				"е не вмер. Ні попівськії тортури, Ні тюремні царські мури, Ані війська муштровані, Ні гар" +
				"мати лаштовані, Ні шпіонське ремесло В гріб його ще не звело.",
			"",
			"",
		},
	)
}
