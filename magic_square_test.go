package cryptolabs_test

import (
	"github.com/xXxRisingTidexXx/cryptography-labs"
	"testing"
)

func TestMagicSquare(t *testing.T) {
	testCipher(
		t,
		cryptolabs.NewMagicSquare(),
		&spec{"", "", ""},
		&spec{"Halloween", "3,2,1", "Halloween"},
		&spec{
			"ВІРТУАЛЬНИЙ КАНАЛ",
			"16,3,2,13,5,10,11,8,9,6,7,12,4,15,14,1",
			"АРІКУИЙЬНАЛ ТНАВ_______________Л",
		},
		&spec{
			"За горами гори, хмарою повиті, Засіяні горем, кровію политі. Споконвіку Прометея Там орел кар" +
				"ає, Що день божий добрі ребра Й серце розбиває. Розбиває, та не вип’є Живущої крові — Вон" +
				"о знову оживає І сміється знову. Не вмирає душа наша, Не вмирає воля. І неситий не виоре " +
				"На дні моря поле. Не скує душ",
			"11,24,7,20,3,4,12,25,8,16,17,5,13,21,9,10,18,1,14,22,23,6,19,2,15",
			"гпар гоом хорои мЗию раа,яоЗ,тінваор,і сіев кр миг рик пСотонопуі.віо ПліюкооТєтер акаяе,м рм" +
				"л Щ ае  Йобньд жі  орийрдба беериєрбрцв,о Реаизбосєва зе.’о їа євввун  ипщ Жкреотин на— о" +
				"соожВвє зиіу Іов  оєямєтв  Неьуизн м.расві  лаєшаНяшми е а,рд вонаувй сеІ  дивинн тио.еНа" +
				"ер  луяу меш е о.єпосн  дркіН",
		},
		&spec{
			"The protagonist of Hamlet is Prince Hamlet of Denmark, son of the recently deceased King Haml" +
				"et, and nephew of King Claudius, his father's brother and successor. Claudius hastily mar" +
				"ried King Hamlet's widow, Gertrude, Hamlet's mother, and took the xx",
			"11,24,7,20,3,4,12,25,8,16,17,5,13,21,9,10,18,1,14,22,23,6,19,2,15",
			"oeoHe ntt opiaagfTsmlr ht nifs HmnetPa ce  mDeroilfyoek,  n r tn oeahtlscreinetcenddamag, Kld" +
				"  aseeH swuepK,  Chidofl niueangrnaeis'dtro srhet   afhhb.sssuc tsudcC ori lhaeusad'rmy  " +
				"srg mKlieHiietaalnrt,aidt' e,ormGe  ulewHwdaxetmonxrootdh, ks e h  t",
		},
		&spec{
			"Вічний революціонер — Дух, що тіло рве до бою, Рве за поступ, щастя й волю, — Він живе, він щ" +
				"е не вмер. Ні попівськії тортури, Ні тюремні царські мури, Ані війська муштровані, Ні гар" +
				"мати лаштовані, Ні шпіонське ремесло В гріб його ще не звело.",
			"16,3,2,13,5,10,11,8,9,6,7,12,4,15,14,1",
			"очіюиворей лніцВірео—, ух Дщ т нР оюв бдое ор ,ла е,атуос ппзщ віят—йю,ол в  В снж щвін ве, и" +
				" енвв оеі  Нр.пміпеНкьиїтуор трі ,скт рр цніемаюьсісм врні Аи, уйіі акнмовтруша ,іьо іааи" +
				" атрмлгтшНьнао,шпі  Нііснві е ео слмеВрргкей  г нщео еовзб_.о____________л",
		},
	)
}
