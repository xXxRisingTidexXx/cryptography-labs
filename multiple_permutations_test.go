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
			"2,4,3,1 4,1,6,3,2,5",
			"И ІДВ АОРЗЗГ ВАЄЖЗ НСАУІ______Ю_ИС_Д____________",
		},
		&spec{
			"ЗАГРОЗА ІСНУЄ ЗАВЖДИ І ВСЮДИ",
			"1,2,3,4 1,2,3,4,5,6",
			"ЗАГРОЗА ІСНУЄ ЗАВЖДИ І ВСЮДИ____________________",
		},
		&spec{
			"За горами гори, хмарою повиті, Засіяні горем, кровію политі. Споконвіку Прометея Там орел кар" +
				"ає, Що день божий добрі ребра Й серце розбиває. Розбиває, та не вип’є Живущої крові — Вон" +
				"о знову оживає І сміється знову. Не вмирає душа наша, Не вмирає воля. І неситий не виоре " +
				"На дні моря поле. Не скує душ",
			"5,4,3,1,6,2 4,1,6,3,2,5",
			"р юапоЗіс яаих рм,мг аоиаогЗр віто,иліто.ивукн ів юопімк ер,іогнр Ско опаєак,рьобнж оле  ряаТ" +
				"ем ремПтоЩд  еор еерц.оРєз а Йрс іеррб йодиб звиоабЖув щиі —вВ в’п єитн  еаиєаб,вїркоо сє" +
				"і тмв .оНув єиІаво ожунз оносз ьня  е,вНо.яв л шааанєудаш  имервиєам рНд  напел .овро еии" +
				"н тей сеІин роіям__________________єудуш Нс  ке______",
		},
		&spec{
			"The protagonist of Hamlet is Prince Hamlet of Denmark, son of the recently deceased King Haml" +
				"et, and nephew of King Claudius, his father's brother and successor. Claudius hastily mar" +
				"ried King Hamlet's widow, Gertrude, Hamlet's mother, and took the xx",
			"5,4,3,1,6,2 4,1,6,3,2,5",
			"Hlm eaiecr nso ifttogonahp Tre  stPis n ooenertcmkrn,aoD  efaelHtm ehf tmtea,lpwee hi gKHnade" +
				"e syedlc a d nnif hasbto hri,sd u alguCfiKon h'rtselduCiai ytmls.rs oscc eurnaed sahus si" +
				"w'd tdureraelHtmKgn  ireiadrwG oe,a d tn _xe_xtreo,ht sem' ma,lHot ohk",
		},
		&spec{
			"Вічний революціонер — Дух, що тіло рве до бою, Рве за поступ, щастя й волю, — Він живе, він щ" +
				"е не вмер. Ні попівськії тортури, Ні тюремні царські мури, Ані війська муштровані, Ні гар" +
				"мати лаштовані, Ні шпіонське ремесло В гріб його ще не звело.",
			"5,4,3,1,6,2 4,1,6,3,2,5",
			" Д ру—і отрлцноюеіров леіинВйч,ощх  оутппс в яойеазв  о ,бРюеодв   са,тщ  еннщ  і.пНев ві,іж " +
				"Виню— л , емервНт  юіаьсцкртирр,уіт коїпсвоьіеінр м шуатм  і,гНіьсвкй ін, А руіимонарівіп" +
				"шНі рем сеа,ів нлтш оартааимнкьоесевзне ______оещг  ій робо Влг о__л_.",
		},
	)
}
