package main

// type relation struct {
// 	Index []struct {
// 		ID              int `json:"id"`
// 		DatesLocations0 struct {
// 			DunedinNewZealand []string `json:"dunedin-new_zealand"`
// 			GeorgiaUsa        []string `json:"georgia-usa"`
// 			LosAngelesUsa     []string `json:"los_angeles-usa"`
// 			NagoyaJapan       []string `json:"nagoya-japan"`
// 			NorthCarolinaUsa  []string `json:"north_carolina-usa"`
// 			OsakaJapan        []string `json:"osaka-japan"`
// 			PenroseNewZealand []string `json:"penrose-new_zealand"`
// 			SaitamaJapan      []string `json:"saitama-japan"`
// 		} `json:"datesLocations,omitempty"`
// 		DatesLocations1 struct {
// 			NoumeaNewCaledonia     []string `json:"noumea-new_caledonia"`
// 			PapeeteFrenchPolynesia []string `json:"papeete-french_polynesia"`
// 			PlayaDelCarmenMexico   []string `json:"playa_del_carmen-mexico"`
// 		} `json:"datesLocations1,omitempty"`
// 		DatesLocations2 struct {
// 			LausanneSwitzerland []string `json:"lausanne-switzerland"`
// 			LondonUk            []string `json:"london-uk"`
// 			LyonFrance          []string `json:"lyon-france"`
// 		} `json:"datesLocations2,omitempty"`
// 		DatesLocations3 struct {
// 			AucklandNewZealand     []string `json:"auckland-new_zealand"`
// 			BratislavaSlovakia     []string `json:"bratislava-slovakia"`
// 			BudapestHungary        []string `json:"budapest-hungary"`
// 			MinskBelarus           []string `json:"minsk-belarus"`
// 			NewSouthWalesAustralia []string `json:"new_south_wales-australia"`
// 			QueenslandAustralia    []string `json:"queensland-australia"`
// 			VictoriaAustralia      []string `json:"victoria-australia"`
// 			YogyakartaIndonesia    []string `json:"yogyakarta-indonesia"`
// 		} `json:"datesLocations3,omitempty"`
// 		DatesLocations4 struct {
// 			CaliforniaUsa []string `json:"california-usa"`
// 			GeorgiaUsa    []string `json:"georgia-usa"`
// 			NevadaUsa     []string `json:"nevada-usa"`
// 		} `json:"datesLocations4,omitempty"`
// 		DatesLocations5 struct {
// 			CaliforniaUsa      []string `json:"california-usa"`
// 			SanIsidroArgentina []string `json:"san_isidro-argentina"`
// 			SaoPauloBrazil     []string `json:"sao_paulo-brazil"`
// 		} `json:"datesLocations5,omitempty"`
// 		DatesLocations6 struct {
// 			ArizonaUsa    []string `json:"arizona-usa"`
// 			CaliforniaUsa []string `json:"california-usa"`
// 			TexasUsa      []string `json:"texas-usa"`
// 		} `json:"datesLocations6,omitempty"`
// 		DatesLocations7 struct {
// 			BilbaoSpain     []string `json:"bilbao-spain"`
// 			BogotaColombia  []string `json:"bogota-colombia"`
// 			GeorgiaUsa      []string `json:"georgia-usa"`
// 			LisbonPortugal  []string `json:"lisbon-portugal"`
// 			SaoPauloBrazil  []string `json:"sao_paulo-brazil"`
// 			StockholmSweden []string `json:"stockholm-sweden"`
// 			WerchterBelgium []string `json:"werchter-belgium"`
// 		} `json:"datesLocations7,omitempty"`
// 		DatesLocations8 struct {
// 			AarhusDenmark     []string `json:"aarhus-denmark"`
// 			DusseldorfGermany []string `json:"dusseldorf-germany"`
// 			ManchesterUk      []string `json:"manchester-uk"`
// 			NewYorkUsa        []string `json:"new_york-usa"`
// 		} `json:"datesLocations8,omitempty"`
// 		DatesLocations9 struct {
// 			BerlinGermany     []string `json:"berlin-germany"`
// 			CopenhagenDenmark []string `json:"copenhagen-denmark"`
// 			FrankfurtGermany  []string `json:"frankfurt-germany"`
// 			StockholmSweden   []string `json:"stockholm-sweden"`
// 			WerchterBelgium   []string `json:"werchter-belgium"`
// 		} `json:"datesLocations9,omitempty"`
// 		DatesLocations10 struct {
// 			CaliforniaUsa []string `json:"california-usa"`
// 			DohaQatar     []string `json:"doha-qatar"`
// 			IllinoisUsa   []string `json:"illinois-usa"`
// 			MinnesotaUsa  []string `json:"minnesota-usa"`
// 			MumbaiIndia   []string `json:"mumbai-india"`
// 		} `json:"datesLocations10,omitempty"`
// 		DatesLocations11 struct {
// 			AbuDhabiUnitedArabEmirates []string `json:"abu_dhabi-united_arab_emirates"`
// 			NewYorkUsa                 []string `json:"new_york-usa"`
// 			PennsylvaniaUsa            []string `json:"pennsylvania-usa"`
// 		} `json:"datesLocations11,omitempty"`
// 		DatesLocations12 struct {
// 			IllinoisUsa      []string `json:"illinois-usa"`
// 			MerkersGermany   []string `json:"merkers-germany"`
// 			WestcliffOnSeaUk []string `json:"westcliff_on_sea-uk"`
// 		} `json:"datesLocations12,omitempty"`
// 		DatesLocations13 struct {
// 			ArizonaUsa    []string `json:"arizona-usa"`
// 			CaliforniaUsa []string `json:"california-usa"`
// 			NevadaUsa     []string `json:"nevada-usa"`
// 		} `json:"datesLocations13,omitempty"`
// 		DatesLocations14 struct {
// 			LondonUk   []string `json:"london-uk"`
// 			MaineUsa   []string `json:"maine-usa"`
// 			NewYorkUsa []string `json:"new_york-usa"`
// 		} `json:"datesLocations14,omitempty"`
// 		DatesLocations15 struct {
// 			AarhusDenmark     []string `json:"aarhus-denmark"`
// 			BerlinGermany     []string `json:"berlin-germany"`
// 			CopenhagenDenmark []string `json:"copenhagen-denmark"`
// 			GothenburgSweden  []string `json:"gothenburg-sweden"`
// 		} `json:"datesLocations15,omitempty"`
// 		DatesLocations16 struct {
// 			FloridaUsa       []string `json:"florida-usa"`
// 			NorthCarolinaUsa []string `json:"north_carolina-usa"`
// 			SouthCarolinaUsa []string `json:"south_carolina-usa"`
// 		} `json:"datesLocations16,omitempty"`
// 		DatesLocations17 struct {
// 			BoulogneBillancourtFrance  []string `json:"boulogne_billancourt-france"`
// 			FreymingMerlebachFrance    []string `json:"freyming_merlebach-france"`
// 			HamburgGermany             []string `json:"hamburg-germany"`
// 			KlagenfurtAustria          []string `json:"klagenfurt-austria"`
// 			NimesFrance                []string `json:"nimes-france"`
// 			OstravaCzechia             []string `json:"ostrava-czechia"`
// 			PagneyDerriereBarineFrance []string `json:"pagney_derriere_barine-france"`
// 			SionSwitzerland            []string `json:"sion-switzerland"`
// 		} `json:"datesLocations17,omitempty"`
// 		DatesLocations18 struct {
// 			LondonUk     []string `json:"london-uk"`
// 			ManchesterUk []string `json:"manchester-uk"`
// 			NevadaUsa    []string `json:"nevada-usa"`
// 		} `json:"datesLocations18,omitempty"`
// 		DatesLocations19 struct {
// 			BarcelonaSpain []string `json:"barcelona-spain"`
// 			MadridSpain    []string `json:"madrid-spain"`
// 			ZaragozaSpain  []string `json:"zaragoza-spain"`
// 		} `json:"datesLocations19,omitempty"`
// 		DatesLocations20 struct {
// 			RecifeBrazil       []string `json:"recife-brazil"`
// 			RioDeJaneiroBrazil []string `json:"rio_de_janeiro-brazil"`
// 		} `json:"datesLocations20,omitempty"`
// 		DatesLocations21 struct {
// 			AmsterdamNetherlands   []string `json:"amsterdam-netherlands"`
// 			BudapestHungary        []string `json:"budapest-hungary"`
// 			BurrianaSpain          []string `json:"burriana-spain"`
// 			CaliforniaUsa          []string `json:"california-usa"`
// 			CuxhavenGermany        []string `json:"cuxhaven-germany"`
// 			LeipzigGermany         []string `json:"leipzig-germany"`
// 			MonchengladbachGermany []string `json:"monchengladbach-germany"`
// 			NapocaRomania          []string `json:"napoca-romania"`
// 			OuluFinland            []string `json:"oulu-finland"`
// 			SalemGermany           []string `json:"salem-germany"`
// 			SkanderborgDenmark     []string `json:"skanderborg-denmark"`
// 		} `json:"datesLocations21,omitempty"`
// 		DatesLocations22 struct {
// 			BerlinGermany      []string `json:"berlin-germany"`
// 			CaliforniaUsa      []string `json:"california-usa"`
// 			CantonUsa          []string `json:"canton-usa"`
// 			DelMarUsa          []string `json:"del_mar-usa"`
// 			LasVegasUsa        []string `json:"las_vegas-usa"`
// 			LisbonPortugal     []string `json:"lisbon-portugal"`
// 			MexicoCityMexico   []string `json:"mexico_city-mexico"`
// 			MonterreyMexico    []string `json:"monterrey-mexico"`
// 			NewYorkUsa         []string `json:"new_york-usa"`
// 			QuebecCanada       []string `json:"quebec-canada"`
// 			RioDeJaneiroBrazil []string `json:"rio_de_janeiro-brazil"`
// 			RiyadhSaudiArabia  []string `json:"riyadh-saudi_arabia"`
// 		} `json:"datesLocations22,omitempty"`
// 		DatesLocations23 struct {
// 			AmsterdamNetherlands   []string `json:"amsterdam-netherlands"`
// 			BirminghamUk           []string `json:"birmingham-uk"`
// 			ChicagoUsa             []string `json:"chicago-usa"`
// 			CopenhagenDenmark      []string `json:"copenhagen-denmark"`
// 			ManchesterUk           []string `json:"manchester-uk"`
// 			MissouriUsa            []string `json:"missouri-usa"`
// 			ParisFrance            []string `json:"paris-france"`
// 			WashingtonUsa          []string `json:"washington-usa"`
// 			WestMelbourneAustralia []string `json:"west_melbourne-australia"`
// 		} `json:"datesLocations23,omitempty"`
// 		DatesLocations24 struct {
// 			BostonUsa              []string `json:"boston-usa"`
// 			CaliforniaUsa          []string `json:"california-usa"`
// 			ClevelandUsa           []string `json:"cleveland-usa"`
// 			MadisonUsa             []string `json:"madison-usa"`
// 			NewYorkUsa             []string `json:"new_york-usa"`
// 			SydneyAustralia        []string `json:"sydney-australia"`
// 			TexasUsa               []string `json:"texas-usa"`
// 			TorontoCanada          []string `json:"toronto-canada"`
// 			UtahUsa                []string `json:"utah-usa"`
// 			WestMelbourneAustralia []string `json:"west_melbourne-australia"`
// 		} `json:"datesLocations24,omitempty"`
// 		DatesLocations25 struct {
// 			AberdeenUk      []string `json:"aberdeen-uk"`
// 			BerlinGermany   []string `json:"berlin-germany"`
// 			CardiffUk       []string `json:"cardiff-uk"`
// 			DublinIreland   []string `json:"dublin-ireland"`
// 			GlasgowUk       []string `json:"glasgow-uk"`
// 			LondonUk        []string `json:"london-uk"`
// 			MadridSpain     []string `json:"madrid-spain"`
// 			ManchesterUk    []string `json:"manchester-uk"`
// 			MilanItaly      []string `json:"milan-italy"`
// 			ParisFrance     []string `json:"paris-france"`
// 			StockholmSweden []string `json:"stockholm-sweden"`
// 			WarsawPoland    []string `json:"warsaw-poland"`
// 		} `json:"datesLocations25,omitempty"`
// 		DatesLocations26 struct {
// 			AmsterdamNetherlands []string `json:"amsterdam-netherlands"`
// 			ColoradoUsa          []string `json:"colorado-usa"`
// 			EindhovenNetherlands []string `json:"eindhoven-netherlands"`
// 			LyonFrance           []string `json:"lyon-france"`
// 			MichiganUsa          []string `json:"michigan-usa"`
// 			NewHampshireUsa      []string `json:"new_hampshire-usa"`
// 			NewYorkUsa           []string `json:"new_york-usa"`
// 			OsloNorway           []string `json:"oslo-norway"`
// 			SochauxFrance        []string `json:"sochaux-france"`
// 			WarsawPoland         []string `json:"warsaw-poland"`
// 		} `json:"datesLocations26,omitempty"`
// 		DatesLocations27 struct {
// 			AalborgDenmark   []string `json:"aalborg-denmark"`
// 			ChangzhouChina   []string `json:"changzhou-china"`
// 			ColoradoUsa      []string `json:"colorado-usa"`
// 			HongKongChina    []string `json:"hong_kong-china"`
// 			HuizhouChina     []string `json:"huizhou-china"`
// 			JakartaIndonesia []string `json:"jakarta-indonesia"`
// 			MichiganUsa      []string `json:"michigan-usa"`
// 			NewYorkUsa       []string `json:"new_york-usa"`
// 			SanyaChina       []string `json:"sanya-china"`
// 			SeattleUsa       []string `json:"seattle-usa"`
// 			TexasUsa         []string `json:"texas-usa"`
// 			TorontoCanada    []string `json:"toronto-canada"`
// 			WashingtonUsa    []string `json:"washington-usa"`
// 		} `json:"datesLocations27,omitempty"`
// 		DatesLocations28 struct {
// 			ColumbiaUsa     []string `json:"columbia-usa"`
// 			GrandRapidsUsa  []string `json:"grand_rapids-usa"`
// 			HersheyUsa      []string `json:"hershey-usa"`
// 			IndianapolisUsa []string `json:"indianapolis-usa"`
// 			KansasCityUsa   []string `json:"kansas_city-usa"`
// 			MontrealUsa     []string `json:"montreal-usa"`
// 			NewarkUsa       []string `json:"newark-usa"`
// 			OmahaUsa        []string `json:"omaha-usa"`
// 			PhiladelphiaUsa []string `json:"philadelphia-usa"`
// 			PittsburghUsa   []string `json:"pittsburgh-usa"`
// 			RosemontUsa     []string `json:"rosemont-usa"`
// 			StLouisUsa      []string `json:"st_louis-usa"`
// 			TorontoUsa      []string `json:"toronto-usa"`
// 			UniondaleUsa    []string `json:"uniondale-usa"`
// 			WashingtonUsa   []string `json:"washington-usa"`
// 		} `json:"datesLocations28,omitempty"`
// 		DatesLocations29 struct {
// 			AtlantaUsa            []string `json:"atlanta-usa"`
// 			FrauenfeldSwitzerland []string `json:"frauenfeld-switzerland"`
// 			HoustonUsa            []string `json:"houston-usa"`
// 			LondonUk              []string `json:"london-uk"`
// 			LosAngelesUsa         []string `json:"los_angeles-usa"`
// 			NewOrleansUsa         []string `json:"new_orleans-usa"`
// 			PhiladelphiaUsa       []string `json:"philadelphia-usa"`
// 			SantiagoChile         []string `json:"santiago-chile"`
// 			SaoPauloBrazil        []string `json:"sao_paulo-brazil"`
// 			TurkuFinland          []string `json:"turku-finland"`
// 		} `json:"datesLocations29,omitempty"`
// 		DatesLocations30 struct {
// 			BostonUsa       []string `json:"boston-usa"`
// 			BrooklynUsa     []string `json:"brooklyn-usa"`
// 			LasVegasUsa     []string `json:"las_vegas-usa"`
// 			MontrealCanada  []string `json:"montreal-canada"`
// 			NewYorkUsa      []string `json:"new_york-usa"`
// 			PhiladelphiaUsa []string `json:"philadelphia-usa"`
// 			TorontoUsa      []string `json:"toronto-usa"`
// 			WashingtonUsa   []string `json:"washington-usa"`
// 		} `json:"datesLocations30,omitempty"`
// 		DatesLocations31 struct {
// 			BerlinGermany     []string `json:"berlin-germany"`
// 			BudapestHungary   []string `json:"budapest-hungary"`
// 			CopenhagenDenmark []string `json:"copenhagen-denmark"`
// 			FrankfurtGermany  []string `json:"frankfurt-germany"`
// 			ImolaItaly        []string `json:"imola-italy"`
// 			KrakowPoland      []string `json:"krakow-poland"`
// 			LondonUk          []string `json:"london-uk"`
// 			StockholmSweden   []string `json:"stockholm-sweden"`
// 			ViennaAustria     []string `json:"vienna-austria"`
// 			ZurichSwitzerland []string `json:"zurich-switzerland"`
// 		} `json:"datesLocations31,omitempty"`
// 		DatesLocations32 struct {
// 			AmityvilleUsa   []string `json:"amityville-usa"`
// 			ChicagoUsa      []string `json:"chicago-usa"`
// 			DetroitUsa      []string `json:"detroit-usa"`
// 			MinneapolisUsa  []string `json:"minneapolis-usa"`
// 			ParisFrance     []string `json:"paris-france"`
// 			PhiladelphiaUsa []string `json:"philadelphia-usa"`
// 		} `json:"datesLocations32,omitempty"`
// 		DatesLocations33 struct {
// 			BerlinGermany []string `json:"berlin-germany"`
// 			CharlotteUsa  []string `json:"charlotte-usa"`
// 			ChicagoUsa    []string `json:"chicago-usa"`
// 			HoustonUsa    []string `json:"houston-usa"`
// 			InglewoodUsa  []string `json:"inglewood-usa"`
// 			LosAngelesUsa []string `json:"los_angeles-usa"`
// 			MadridSpain   []string `json:"madrid-spain"`
// 			OaklandUsa    []string `json:"oakland-usa"`
// 		} `json:"datesLocations33,omitempty"`
// 		DatesLocations34 struct {
// 			AnaheimUsa    []string `json:"anaheim-usa"`
// 			BirminghamUk  []string `json:"birmingham-uk"`
// 			BrooklynUsa   []string `json:"brooklyn-usa"`
// 			ChicagoUsa    []string `json:"chicago-usa"`
// 			CincinnatiUsa []string `json:"cincinnati-usa"`
// 			WindsorCanada []string `json:"windsor-canada"`
// 		} `json:"datesLocations34,omitempty"`
// 		DatesLocations35 struct {
// 			AucklandNewZealand []string `json:"auckland-new_zealand"`
// 			BrisbaneAustralia  []string `json:"brisbane-australia"`
// 			ManilaPhilippines  []string `json:"manila-philippines"`
// 			MelbourneAustralia []string `json:"melbourne-australia"`
// 			MumbaiIndia        []string `json:"mumbai-india"`
// 			SydneyAustralia    []string `json:"sydney-australia"`
// 		} `json:"datesLocations35,omitempty"`
// 		DatesLocations36 struct {
// 			BogotaColombia     []string `json:"bogota-colombia"`
// 			LimaPeru           []string `json:"lima-peru"`
// 			RioDeJaneiroBrazil []string `json:"rio_de_janeiro-brazil"`
// 			SanIsidroArgentina []string `json:"san_isidro-argentina"`
// 			SantiagoChile      []string `json:"santiago-chile"`
// 			SaoPauloBrazil     []string `json:"sao_paulo-brazil"`
// 		} `json:"datesLocations36,omitempty"`
// 		DatesLocations37 struct {
// 			AntwerpBelgium       []string `json:"antwerp-belgium"`
// 			GlasgowUk            []string `json:"glasgow-uk"`
// 			GroningenNetherlands []string `json:"groningen-netherlands"`
// 			LondonUk             []string `json:"london-uk"`
// 			ParisFrance          []string `json:"paris-france"`
// 			ViennaAustria        []string `json:"vienna-austria"`
// 		} `json:"datesLocations37,omitempty"`
// 		DatesLocations38 struct {
// 			BostonUsa        []string `json:"boston-usa"`
// 			ChicagoUsa       []string `json:"chicago-usa"`
// 			MexicoCityMexico []string `json:"mexico_city-mexico"`
// 			PhiladelphiaUsa  []string `json:"philadelphia-usa"`
// 			PicoRiveraUsa    []string `json:"pico_rivera-usa"`
// 		} `json:"datesLocations38,omitempty"`
// 		DatesLocations39 struct {
// 			BerwynUsa  []string `json:"berwyn-usa"`
// 			DallasUsa  []string `json:"dallas-usa"`
// 			GeorgiaUsa []string `json:"georgia-usa"`
// 			HoustonUsa []string `json:"houston-usa"`
// 			LondonUk   []string `json:"london-uk"`
// 			NewYorkUsa []string `json:"new_york-usa"`
// 		} `json:"datesLocations39,omitempty"`
// 		DatesLocations40 struct {
// 			BirminghamUk     []string `json:"birmingham-uk"`
// 			BrixtonUk        []string `json:"brixton-uk"`
// 			CaliforniaUsa    []string `json:"california-usa"`
// 			LondonUk         []string `json:"london-uk"`
// 			RotselaarBelgium []string `json:"rotselaar-belgium"`
// 		} `json:"datesLocations40,omitempty"`
// 		DatesLocations41 struct {
// 			AlabamaUsa           []string `json:"alabama-usa"`
// 			AthensGreece         []string `json:"athens-greece"`
// 			CaliforniaUsa        []string `json:"california-usa"`
// 			FlorenceItaly        []string `json:"florence-italy"`
// 			LandgraafNetherlands []string `json:"landgraaf-netherlands"`
// 			LosAngelesUsa        []string `json:"los_angeles-usa"`
// 			MassachusettsUsa     []string `json:"massachusetts-usa"`
// 			RioDeJaneiroBrazil   []string `json:"rio_de_janeiro-brazil"`
// 		} `json:"datesLocations41,omitempty"`
// 		DatesLocations42 struct {
// 			AbuDhabiUnitedArabEmirates []string `json:"abu_dhabi-united_arab_emirates"`
// 			BurswoodAustralia          []string `json:"burswood-australia"`
// 			MelbourneAustralia         []string `json:"melbourne-australia"`
// 			SydneyAustralia            []string `json:"sydney-australia"`
// 			WellingtonNewZealand       []string `json:"wellington-new_zealand"`
// 		} `json:"datesLocations42,omitempty"`
// 		DatesLocations43 struct {
// 			BangkokThailand   []string `json:"bangkok-thailand"`
// 			HongKongChina     []string `json:"hong_kong-china"`
// 			LosAngelesUsa     []string `json:"los_angeles-usa"`
// 			MadridSpain       []string `json:"madrid-spain"`
// 			ManilaPhilippines []string `json:"manila-philippines"`
// 			SeoulSouthKorea   []string `json:"seoul-south_korea"`
// 			SevilleSpain      []string `json:"seville-spain"`
// 			TaipeiTaiwan      []string `json:"taipei-taiwan"`
// 		} `json:"datesLocations43,omitempty"`
// 		DatesLocations44 struct {
// 			BeloHorizonteBrazil  []string `json:"belo_horizonte-brazil"`
// 			BuenosAiresArgentina []string `json:"buenos_aires-argentina"`
// 			MannheimGermany      []string `json:"mannheim-germany"`
// 			MunichGermany        []string `json:"munich-germany"`
// 			PortoAlegreBrazil    []string `json:"porto_alegre-brazil"`
// 			SanFranciscoUsa      []string `json:"san_francisco-usa"`
// 			SantiagoChile        []string `json:"santiago-chile"`
// 			SaoPauloBrazil       []string `json:"sao_paulo-brazil"`
// 		} `json:"datesLocations44,omitempty"`
// 		DatesLocations45 struct {
// 			CaliforniaUsa     []string `json:"california-usa"`
// 			LaPlataArgentina  []string `json:"la_plata-argentina"`
// 			LondonUk          []string `json:"london-uk"`
// 			PortoAlegreBrazil []string `json:"porto_alegre-brazil"`
// 			SaoPauloBrazil    []string `json:"sao_paulo-brazil"`
// 		} `json:"datesLocations45,omitempty"`
// 		DatesLocations46 struct {
// 			BrasiliaBrazil                []string `json:"brasilia-brazil"`
// 			DubaiUnitedArabEmirates       []string `json:"dubai-united_arab_emirates"`
// 			FloridaUsa                    []string `json:"florida-usa"`
// 			MexicoCityMexico              []string `json:"mexico_city-mexico"`
// 			SantiagoChile                 []string `json:"santiago-chile"`
// 			SaoPauloBrazil                []string `json:"sao_paulo-brazil"`
// 			WillemstadNetherlandsAntilles []string `json:"willemstad-netherlands_antilles"`
// 		} `json:"datesLocations46,omitempty"`
// 		DatesLocations47 struct {
// 			ArrasFrance         []string `json:"arras-france"`
// 			CaliforniaUsa       []string `json:"california-usa"`
// 			GdyniaPoland        []string `json:"gdynia-poland"`
// 			IllinoisUsa         []string `json:"illinois-usa"`
// 			OklahomaUsa         []string `json:"oklahoma-usa"`
// 			ScheesselGermany    []string `json:"scheessel-germany"`
// 			StGallenSwitzerland []string `json:"st_gallen-switzerland"`
// 			TexasUsa            []string `json:"texas-usa"`
// 		} `json:"datesLocations47,omitempty"`
// 		DatesLocations48 struct {
// 			ArizonaUsa    []string `json:"arizona-usa"`
// 			CaliforniaUsa []string `json:"california-usa"`
// 			FloridaUsa    []string `json:"florida-usa"`
// 			WashingtonUsa []string `json:"washington-usa"`
// 		} `json:"datesLocations48,omitempty"`
// 		DatesLocations49 struct {
// 			BuenosAiresArgentina []string `json:"buenos_aires-argentina"`
// 			LimaPeru             []string `json:"lima-peru"`
// 			MexicoCityMexico     []string `json:"mexico_city-mexico"`
// 			SantiagoChile        []string `json:"santiago-chile"`
// 			SaoPauloBrazil       []string `json:"sao_paulo-brazil"`
// 		} `json:"datesLocations49,omitempty"`
// 		DatesLocations50 struct {
// 			BogotaColombia     []string `json:"bogota-colombia"`
// 			LisbonPortugal     []string `json:"lisbon-portugal"`
// 			MassachusettsUsa   []string `json:"massachusetts-usa"`
// 			MilanItaly         []string `json:"milan-italy"`
// 			NevadaUsa          []string `json:"nevada-usa"`
// 			NickelsdorfAustria []string `json:"nickelsdorf-austria"`
// 			SanJoseCostaRica   []string `json:"san_jose-costa_rica"`
// 		} `json:"datesLocations50,omitempty"`
// 		DatesLocations51 struct {
// 			ColoradoUsa     []string `json:"colorado-usa"`
// 			MilanItaly      []string `json:"milan-italy"`
// 			MunichGermany   []string `json:"munich-germany"`
// 			NevadaUsa       []string `json:"nevada-usa"`
// 			OregonUsa       []string `json:"oregon-usa"`
// 			PragueCzechia   []string `json:"prague-czechia"`
// 			VancouverCanada []string `json:"vancouver-canada"`
// 		} `json:"datesLocations51,omitempty"`
// 	} `json:"index"`
// }