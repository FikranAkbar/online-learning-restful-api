package seeder

import (
	gormseeder "github.com/kachit/gorm-seeder"
	"gorm.io/gorm"
	"online-learning-restful-api/app/database/entity"
)

type MasterCitySeeder struct {
	gormseeder.SeederAbstract
}

func NewMasterCitySeeder(cfg gormseeder.SeederConfiguration) *MasterCitySeeder {
	return &MasterCitySeeder{gormseeder.SeederAbstract{Configuration: cfg}}
}

func InitCityData() []entity.MasterCity {
	cityStrings := []string{
		"PIDIE JAYA", "SIMEULUE", "BIREUEN", "ACEH TIMUR", "ACEH UTARA", "PIDIE", "ACEH BARAT DAYA", "GAYO LUES", "ACEH SELATAN", "ACEH TAMIANG", "ACEH BESAR", "ACEH TENGGARA", "BENER MERIAH", "ACEH JAYA", "LHOKSEUMAWE", "ACEH BARAT", "NAGAN RAYA", "LANGSA", "BANDA ACEH", "ACEH SINGKIL", "SABANG", "ACEH TENGAH", "SUBULUSSALAM", "NIAS SELATAN", "MANDAILING NATAL", "DAIRI", "LABUHAN BATU UTARA", "TAPANULI UTARA", "SIMALUNGUN", "LANGKAT", "SERDANG BEDAGAI", "TAPANULI SELATAN", "ASAHAN", "PADANG LAWAS UTARA", "PADANG LAWAS", "LABUHAN BATU SELATAN", "PADANG SIDEMPUAN", "TOBA SAMOSIR", "TAPANULI TENGAH", "HUMBANG HASUNDUTAN", "SIBOLGA", "BATU BARA", "SAMOSIR", "PEMATANG SIANTAR", "LABUHAN BATU", "DELI SERDANG", "GUNUNGSITOLI", "NIAS UTARA", "NIAS", "KARO", "NIAS BARAT", "MEDAN", "PAKPAK BHARAT", "TEBING TINGGI", "BINJAI", "TANJUNG BALAI", "DHARMASRAYA", "SOLOK SELATAN", "SIJUNJUNG (SAWAH LUNTO SIJUNJUNG)", "PASAMAN BARAT", "SOLOK", "PASAMAN", "PARIAMAN", "TANAH DATAR", "PADANG PARIAMAN", "PESISIR SELATAN", "PADANG", "SAWAH LUNTO", "LIMA PULUH KOTO / KOTA", "AGAM", "PAYAKUMBUH", "BUKITTINGGI", "PADANG PANJANG", "KEPULAUAN MENTAWAI", "INDRAGIRI HILIR", "KUANTAN SINGINGI", "PELALAWAN", "PEKANBARU", "ROKAN HILIR", "BENGKALIS", "INDRAGIRI HULU", "ROKAN HULU", "KAMPAR", "KEPULAUAN MERANTI", "DUMAI", "SIAK", "TEBO", "TANJUNG JABUNG BARAT", "MUARO JAMBI", "KERINCI", "MERANGIN", "BUNGO", "TANJUNG JABUNG TIMUR", "SUNGAIPENUH", "BATANG HARI", "JAMBI", "SAROLANGUN", "PALEMBANG", "LAHAT", "OGAN KOMERING ULU TIMUR", "MUSI BANYUASIN", "PAGAR ALAM", "OGAN KOMERING ULU SELATAN", "BANYUASIN", "MUSI RAWAS", "MUARA ENIM", "OGAN KOMERING ULU", "OGAN KOMERING ILIR", "EMPAT LAWANG", "LUBUK LINGGAU", "PRABUMULIH", "OGAN ILIR", "BENGKULU TENGAH", "REJANG LEBONG", "MUKO MUKO", "KAUR", "BENGKULU UTARA", "LEBONG", "KEPAHIANG", "BENGKULU SELATAN", "SELUMA", "BENGKULU", "LAMPUNG UTARA", "WAY KANAN", "LAMPUNG TENGAH", "MESUJI", "PRINGSEWU", "LAMPUNG TIMUR", "LAMPUNG SELATAN", "TULANG BAWANG", "TULANG BAWANG BARAT", "TANGGAMUS", "LAMPUNG BARAT", "PESISIR BARAT", "PESAWARAN", "BANDAR LAMPUNG", "METRO", "BELITUNG", "BELITUNG TIMUR", "BANGKA", "BANGKA SELATAN", "BANGKA BARAT", "PANGKAL PINANG", "BANGKA TENGAH", "KEPULAUAN ANAMBAS", "BINTAN", "NATUNA", "BATAM", "TANJUNG PINANG", "KARIMUN", "LINGGA", "JAKARTA UTARA", "JAKARTA BARAT", "JAKARTA TIMUR", "JAKARTA SELATAN", "JAKARTA PUSAT", "KEPULAUAN SERIBU", "DEPOK", "KARAWANG", "CIREBON", "BANDUNG", "SUKABUMI", "SUMEDANG", "INDRAMAYU", "MAJALENGKA", "KUNINGAN", "TASIKMALAYA", "CIAMIS", "SUBANG", "PURWAKARTA", "BOGOR", "BEKASI", "GARUT", "PANGANDARAN", "CIANJUR", "BANJAR", "BANDUNG BARAT", "CIMAHI", "PURBALINGGA", "KEBUMEN", "MAGELANG", "CILACAP", "BATANG", "BANJARNEGARA", "BLORA", "BREBES", "BANYUMAS", "WONOSOBO", "TEGAL", "PURWOREJO", "PATI", "SUKOHARJO", "KARANGANYAR", "PEKALONGAN", "PEMALANG", "BOYOLALI", "GROBOGAN", "SEMARANG", "DEMAK", "REMBANG", "KLATEN", "KUDUS", "TEMANGGUNG", "SRAGEN", "JEPARA", "WONOGIRI", "KENDAL", "SURAKARTA (SOLO)", "SALATIGA", "SLEMAN", "BANTUL", "YOGYAKARTA", "GUNUNG KIDUL", "KULON PROGO", "GRESIK", "KEDIRI", "SAMPANG", "BANGKALAN", "SUMENEP", "SITUBONDO", "SURABAYA", "JEMBER", "PAMEKASAN", "JOMBANG", "PROBOLINGGO", "BANYUWANGI", "PASURUAN", "BOJONEGORO", "BONDOWOSO", "MAGETAN", "LUMAJANG", "MALANG", "BLITAR", "SIDOARJO", "LAMONGAN", "PACITAN", "TULUNGAGUNG", "MOJOKERTO", "MADIUN", "PONOROGO", "NGAWI", "NGANJUK", "TUBAN", "TRENGGALEK", "BATU", "TANGERANG", "SERANG", "PANDEGLANG", "LEBAK", "TANGERANG SELATAN", "CILEGON", "KLUNGKUNG", "KARANGASEM", "BANGLI", "TABANAN", "GIANYAR", "BADUNG", "JEMBRANA", "BULELENG", "DENPASAR", "MATARAM", "DOMPU", "SUMBAWA BARAT", "SUMBAWA", "LOMBOK TENGAH", "LOMBOK TIMUR", "LOMBOK UTARA", "LOMBOK BARAT", "BIMA", "TIMOR TENGAH SELATAN", "FLORES TIMUR", "ALOR", "ENDE", "NAGEKEO", "KUPANG", "SIKKA", "NGADA", "TIMOR TENGAH UTARA", "BELU", "LEMBATA", "SUMBA BARAT DAYA", "SUMBA BARAT", "SUMBA TENGAH", "SUMBA TIMUR", "ROTE NDAO", "MANGGARAI TIMUR", "MANGGARAI", "SABU RAIJUA", "MANGGARAI BARAT", "LANDAK", "KETAPANG", "SINTANG", "KUBU RAYA", "PONTIANAK", "KAYONG UTARA", "BENGKAYANG", "KAPUAS HULU", "SAMBAS", "SINGKAWANG", "SANGGAU", "MELAWI", "SEKADAU", "KOTAWARINGIN TIMUR", "SUKAMARA", "KOTAWARINGIN BARAT", "BARITO TIMUR", "KAPUAS", "PULANG PISAU", "LAMANDAU", "SERUYAN", "KATINGAN", "BARITO SELATAN", "MURUNG RAYA", "BARITO UTARA", "GUNUNG MAS", "PALANGKA RAYA", "TAPIN", "BANJAR", "HULU SUNGAI TENGAH", "TABALONG", "HULU SUNGAI UTARA", "BALANGAN", "TANAH BUMBU", "BANJARMASIN", "KOTABARU", "TANAH LAUT", "HULU SUNGAI SELATAN", "BARITO KUALA", "BANJARBARU", "KUTAI BARAT", "SAMARINDA", "PASER", "KUTAI KARTANEGARA", "BERAU", "PENAJAM PASER UTARA", "BONTANG", "KUTAI TIMUR", "BALIKPAPAN", "MALINAU", "NUNUKAN", "BULUNGAN (BULONGAN)", "TANA TIDUNG", "TARAKAN", "BOLAANG MONGONDOW (BOLMONG)", "BOLAANG MONGONDOW SELATAN", "MINAHASA SELATAN", "BITUNG", "MINAHASA", "KEPULAUAN SANGIHE", "MINAHASA UTARA", "KEPULAUAN TALAUD", "KEPULAUAN SIAU TAGULANDANG BIARO (SITARO)", "MANADO", "BOLAANG MONGONDOW UTARA", "BOLAANG MONGONDOW TIMUR", "MINAHASA TENGGARA", "KOTAMOBAGU", "TOMOHON", "BANGGAI KEPULAUAN", "TOLI-TOLI", "PARIGI MOUTONG", "BUOL", "DONGGALA", "POSO", "MOROWALI", "TOJO UNA-UNA", "BANGGAI", "SIGI", "PALU", "MAROS", "WAJO", "BONE", "SOPPENG", "SIDENRENG RAPPANG / RAPANG", "TAKALAR", "BARRU", "LUWU TIMUR", "SINJAI", "PANGKAJENE KEPULAUAN", "PINRANG", "JENEPONTO", "PALOPO", "TORAJA UTARA", "LUWU", "BULUKUMBA", "MAKASSAR", "SELAYAR (KEPULAUAN SELAYAR)", "TANA TORAJA", "LUWU UTARA", "BANTAENG", "GOWA", "ENREKANG", "PAREPARE", "KOLAKA", "MUNA", "KONAWE SELATAN", "KENDARI", "KONAWE", "KONAWE UTARA", "KOLAKA UTARA", "BUTON", "BOMBANA", "WAKATOBI", "BAU-BAU", "BUTON UTARA", "GORONTALO UTARA", "BONE BOLANGO", "GORONTALO", "BOALEMO", "POHUWATO", "MAJENE", "MAMUJU", "MAMUJU UTARA", "POLEWALI MANDAR", "MAMASA", "MALUKU TENGGARA BARAT", "MALUKU TENGGARA", "SERAM BAGIAN BARAT", "MALUKU TENGAH", "SERAM BAGIAN TIMUR", "MALUKU BARAT DAYA", "AMBON", "BURU", "BURU SELATAN", "KEPULAUAN ARU", "TUAL", "HALMAHERA BARAT", "TIDORE KEPULAUAN", "TERNATE", "PULAU MOROTAI", "KEPULAUAN SULA", "HALMAHERA SELATAN", "HALMAHERA TENGAH", "HALMAHERA TIMUR", "HALMAHERA UTARA", "YALIMO", "DOGIYAI", "ASMAT", "JAYAPURA", "PANIAI", "MAPPI", "TOLIKARA", "PUNCAK JAYA", "PEGUNUNGAN BINTANG", "JAYAWIJAYA", "LANNY JAYA", "NDUGA", "BIAK NUMFOR", "KEPULAUAN YAPEN (YAPEN WAROPEN)", "PUNCAK", "INTAN JAYA", "WAROPEN", "NABIRE", "MIMIKA", "BOVEN DIGOEL", "YAHUKIMO", "SARMI", "MERAUKE", "DEIYAI (DELIYAI)", "KEEROM", "SUPIORI", "MAMBERAMO RAYA", "MAMBERAMO TENGAH", "RAJA AMPAT", "MANOKWARI SELATAN", "MANOKWARI", "KAIMANA", "MAYBRAT", "SORONG SELATAN", "FAKFAK", "PEGUNUNGAN ARFAK", "TAMBRAUW", "SORONG", "TELUK WONDAMA", "TELUK BINTUNI",
	}
	provinceIds := []int{
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 6, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 8, 9, 9, 9, 9, 9, 9, 9, 10, 10, 10, 10, 10, 10, 10, 11, 11, 11, 11, 11, 11, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 12, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 13, 14, 14, 14, 14, 14, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 16, 16, 16, 16, 16, 16, 17, 17, 17, 17, 17, 17, 17, 17, 17, 18, 18, 18, 18, 18, 18, 18, 18, 18, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 19, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 20, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 21, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 23, 23, 23, 23, 23, 23, 23, 23, 23, 24, 24, 24, 24, 24, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 25, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 27, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 28, 29, 29, 29, 29, 29, 30, 30, 30, 30, 30, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 32, 32, 32, 32, 32, 32, 32, 32, 32, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 33, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34,
	}

	var cities []entity.MasterCity

	for i := 1; i <= len(cityStrings); i++ {
		cities = append(cities, entity.MasterCity{
			Model:      gorm.Model{ID: uint(i)},
			CityName:   cityStrings[i-1],
			ProvinceId: uint(provinceIds[i-1]),
		})
	}

	return cities
}

func (s *MasterCitySeeder) Seed(db *gorm.DB) error {
	data := InitCityData()
	return db.CreateInBatches(data, len(data)).Error
}

func (s *MasterCitySeeder) Clear(db *gorm.DB) error {
	ent := &entity.MasterCity{}
	return s.SeederAbstract.Delete(db, ent.TableName())
}