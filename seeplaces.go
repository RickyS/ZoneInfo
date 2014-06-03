package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now() // now is type time.Time, not a pointer.
	stdform := "Mon Jan 2 15:04:05 -0700 MST 2006"

	placeCnt := 0
	badCnt := 0
	for _, line := range places {
		placeCnt++

		pLoc, err := time.LoadLocation(line) //
		if nil != err {
			fmt.Printf("ERR at Location: '%v'\n", err)
			badCnt++
		} else {
			//fmt.Printf("ok: %s\n", pLoc)
			display := now.In(pLoc).Format(stdform)
			fmt.Printf(" %34s: %s\n", display, pLoc)
		}

	}
	fmt.Printf("[%3d Bad places out of %3d]\n", badCnt, placeCnt)

}

// 99% of this list was generated by small script in makenamelist.sh.
var places = [...]string{
	"Africa/Lubumbashi",
	"Africa/Windhoek",
	"Africa/Niamey",
	"Africa/Ouagadougou",
	"Africa/Brazzaville",
	"Africa/Algiers",
	"Africa/Lagos",
	"Africa/Juba",
	"Africa/Djibouti",
	"Africa/Tunis",
	"Africa/Dakar",
	"Africa/Malabo",
	"Africa/Mbabane",
	"Africa/Dar_es_Salaam",
	"Africa/Conakry",
	"Africa/Bamako",
	"Africa/Libreville",
	"Africa/Mogadishu",
	"Africa/Blantyre",
	"Africa/Kigali",
	"Africa/Ndjamena",
	"Africa/Accra",
	"Africa/Nairobi",
	"Africa/Lusaka",
	"Africa/Kampala",
	"Africa/Casablanca",
	"Africa/Freetown",
	"Africa/Ceuta",
	"Africa/Banjul",
	"Africa/Gaborone",
	"Africa/Asmera",
	"Africa/Sao_Tome",
	"Africa/Timbuktu",
	"Africa/Nouakchott",
	"Africa/Maputo",
	"Africa/Kinshasa",
	"Africa/Harare",
	"Africa/Bissau",
	"Africa/Monrovia",
	"Africa/Maseru",
	"Africa/Porto-Novo",
	"Africa/Cairo",
	"Africa/Asmara",
	"Africa/Lome",
	"Africa/Douala",
	"Africa/Tripoli",
	"Africa/Addis_Ababa",
	"Africa/Bangui",
	"Africa/El_Aaiun",
	"Africa/Bujumbura",
	"Africa/Johannesburg",
	"Africa/Abidjan",
	"Africa/Khartoum",
	"Africa/Luanda",
	"America/Vancouver",
	"America/Creston",
	"America/Manaus",
	"America/Scoresbysund",
	"America/Grand_Turk",
	"America/Porto_Acre",
	"America/Shiprock",
	"America/Lower_Princes",
	"America/Monterrey",
	"America/Recife",
	"America/Dawson_Creek",
	"America/Porto_Velho",
	"America/Metlakatla",
	"America/Anchorage",
	"America/Marigot",
	"America/Denver",
	"America/Mexico_City",
	"America/Dawson",
	"America/Whitehorse",
	"America/Ojinaga",
	"America/Knox_IN",
	"America/Goose_Bay",
	"America/Rainy_River",
	"America/Moncton",
	"America/Regina",
	"America/Santa_Isabel",
	"America/North_Dakota/Center",
	"America/North_Dakota/Beulah",
	"America/North_Dakota/New_Salem",
	"America/Cayenne",
	"America/Curacao",
	"America/Godthab",
	"America/Jamaica",
	"America/Boa_Vista",
	"America/Costa_Rica",
	"America/Montserrat",
	"America/Guadeloupe",
	"America/La_Paz",
	"America/Campo_Grande",
	"America/Indianapolis",
	"America/Port-au-Prince",
	"America/Martinique",
	"America/Cancun",
	"America/Montreal",
	"America/Antigua",
	"America/St_Kitts",
	"America/Nome",
	"America/Sao_Paulo",
	"America/Boise",
	"America/Edmonton",
	"America/Tijuana",
	"America/Jujuy",
	"America/Chicago",
	"America/Rio_Branco",
	"America/Cordoba",
	"America/Cayman",
	"America/Atka",
	"America/Menominee",
	"America/Santiago",
	"America/Santarem",
	"America/Merida",
	"America/Aruba",
	"America/Panama",
	"America/Puerto_Rico",
	"America/Dominica",
	"America/Sitka",
	"America/New_York",
	"America/Halifax",
	"America/Cuiaba",
	"America/Iqaluit",
	"America/Winnipeg",
	"America/Coral_Harbour",
	"America/Tortola",
	"America/Catamarca",
	"America/Bahia_Banderas",
	"America/Rankin_Inlet",
	"America/Danmarkshavn",
	"America/Belem",
	"America/Swift_Current",
	"America/Nipigon",
	"America/Phoenix",
	"America/Port_of_Spain",
	"America/Toronto",
	"America/St_Thomas",
	"America/Argentina/Salta",
	"America/Argentina/ComodRivadavia",
	"America/Argentina/San_Luis",
	"America/Argentina/Ushuaia",
	"America/Argentina/Jujuy",
	"America/Argentina/Cordoba",
	"America/Argentina/Catamarca",
	"America/Argentina/Buenos_Aires",
	"America/Argentina/Mendoza",
	"America/Argentina/Tucuman",
	"America/Argentina/Rio_Gallegos",
	"America/Argentina/La_Rioja",
	"America/Argentina/San_Juan",
	"America/Inuvik",
	"America/Guatemala",
	"America/Santo_Domingo",
	"America/Buenos_Aires",
	"America/Thule",
	"America/Pangnirtung",
	"America/Hermosillo",
	"America/Resolute",
	"America/Virgin",
	"America/Bahia",
	"America/Mendoza",
	"America/Kralendijk",
	"America/Blanc-Sablon",
	"America/Matamoros",
	"America/Asuncion",
	"America/St_Barthelemy",
	"America/Yakutat",
	"America/Guyana",
	"America/Indiana/Indianapolis",
	"America/Indiana/Vevay",
	"America/Indiana/Marengo",
	"America/Indiana/Vincennes",
	"America/Indiana/Petersburg",
	"America/Indiana/Knox",
	"America/Indiana/Tell_City",
	"America/Indiana/Winamac",
	"America/Nassau",
	"America/Ensenada",
	"America/Chihuahua",
	"America/Lima",
	"America/St_Lucia",
	"America/Belize",
	"America/Detroit",
	"America/Cambridge_Bay",
	"America/Glace_Bay",
	"America/Adak",
	"America/Noronha",
	"America/Bogota",
	"America/Araguaina",
	"America/Juneau",
	"America/Montevideo",
	"America/Managua",
	"America/Thunder_Bay",
	"America/Louisville",
	"America/El_Salvador",
	"America/St_Johns",
	"America/Paramaribo",
	"America/Los_Angeles",
	"America/Barbados",
	"America/Havana",
	"America/Anguilla",
	"America/Yellowknife",
	"America/Guayaquil",
	"America/Grenada",
	"America/Tegucigalpa",
	"America/Maceio",
	"America/Rosario",
	"America/Mazatlan",
	"America/Caracas",
	"America/Eirunepe",
	"America/Atikokan",
	"America/Kentucky/Monticello",
	"America/Kentucky/Louisville",
	"America/Fort_Wayne",
	"America/Fortaleza",
	"America/St_Vincent",
	"America/Miquelon",
	"Antarctica/Vostok",
	"Antarctica/South_Pole",
	"Antarctica/Macquarie",
	"Antarctica/Rothera",
	"Antarctica/Syowa",
	"Antarctica/DumontDUrville",
	"Antarctica/Palmer",
	"Antarctica/Casey",
	"Antarctica/Davis",
	"Antarctica/McMurdo",
	"Antarctica/Mawson",
	"Arctic/Longyearbyen",
	"Asia/Amman",
	"Asia/Phnom_Penh",
	"Asia/Yakutsk",
	"Asia/Yerevan",
	"Asia/Manila",
	"Asia/Ust-Nera",
	"Asia/Sakhalin",
	"Asia/Tel_Aviv",
	"Asia/Seoul",
	"Asia/Aden",
	"Asia/Taipei",
	"Asia/Jerusalem",
	"Asia/Qatar",
	"Asia/Kuwait",
	"Asia/Ashkhabad",
	"Asia/Kolkata",
	"Asia/Ashgabat",
	"Asia/Tokyo",
	"Asia/Tehran",
	"Asia/Makassar",
	"Asia/Macao",
	"Asia/Macau",
	"Asia/Brunei",
	"Asia/Dili",
	"Asia/Shanghai",
	"Asia/Jayapura",
	"Asia/Colombo",
	"Asia/Istanbul",
	"Asia/Dhaka",
	"Asia/Riyadh87",
	"Asia/Yekaterinburg",
	"Asia/Riyadh88",
	"Asia/Harbin",
	"Asia/Karachi",
	"Asia/Thimphu",
	"Asia/Krasnoyarsk",
	"Asia/Almaty",
	"Asia/Katmandu",
	"Asia/Pontianak",
	"Asia/Magadan",
	"Asia/Chongqing",
	"Asia/Rangoon",
	"Asia/Kuala_Lumpur",
	"Asia/Irkutsk",
	"Asia/Bangkok",
	"Asia/Kamchatka",
	"Asia/Bishkek",
	"Asia/Pyongyang",
	"Asia/Beirut",
	"Asia/Hong_Kong",
	"Asia/Jakarta",
	"Asia/Hovd",
	"Asia/Kathmandu",
	"Asia/Dubai",
	"Asia/Saigon",
	"Asia/Aqtobe",
	"Asia/Ulan_Bator",
	"Asia/Gaza",
	"Asia/Qyzylorda",
	"Asia/Tbilisi",
	"Asia/Novosibirsk",
	"Asia/Muscat",
	"Asia/Omsk",
	"Asia/Calcutta",
	"Asia/Khandyga",
	"Asia/Bahrain",
	"Asia/Chungking",
	"Asia/Samarkand",
	"Asia/Oral",
	"Asia/Novokuznetsk",
	"Asia/Singapore",
	"Asia/Riyadh",
	"Asia/Kashgar",
	"Asia/Vientiane",
	"Asia/Tashkent",
	"Asia/Baghdad",
	"Asia/Ulaanbaatar",
	"Asia/Nicosia",
	"Asia/Thimbu",
	"Asia/Choibalsan",
	"Asia/Aqtau",
	"Asia/Ujung_Pandang",
	"Asia/Urumqi",
	"Asia/Vladivostok",
	"Asia/Hebron",
	"Asia/Ho_Chi_Minh",
	"Asia/Dushanbe",
	"Asia/Kabul",
	"Asia/Anadyr",
	"Asia/Baku",
	"Asia/Riyadh89",
	"Asia/Damascus",
	"Asia/Kuching",
	"Asia/Dacca",
	"Atlantic/South_Georgia",
	"Atlantic/Jan_Mayen",
	"Atlantic/Madeira",
	"Atlantic/Stanley",
	"Atlantic/Faeroe",
	"Atlantic/Canary",
	"Atlantic/Faroe",
	"Atlantic/Azores",
	"Atlantic/Reykjavik",
	"Atlantic/St_Helena",
	"Atlantic/Cape_Verde",
	"Atlantic/Bermuda",
	"Australia/Lord_Howe",
	"Australia/Adelaide",
	"Australia/Yancowinna",
	"Australia/NSW",
	"Australia/Sydney",
	"Australia/ACT",
	"Australia/Tasmania",
	"Australia/North",
	"Australia/Brisbane",
	"Australia/Currie",
	"Australia/Broken_Hill",
	"Australia/Eucla",
	"Australia/LHI",
	"Australia/Hobart",
	"Australia/Darwin",
	"Australia/Queensland",
	"Australia/West",
	"Australia/Melbourne",
	"Australia/Canberra",
	"Australia/Perth",
	"Australia/Lindeman",
	"Australia/South",
	"Australia/Victoria",
	"Brazil/DeNoronha",
	"Brazil/Acre",
	"Brazil/West",
	"Brazil/East",
	"CET",
	"CST6CDT",
	"Canada/Newfoundland",
	"Canada/Central",
	"Canada/Yukon",
	"Canada/Mountain",
	"Canada/East-Saskatchewan",
	"Canada/Pacific",
	"Canada/Saskatchewan",
	"Canada/Eastern",
	"Canada/Atlantic",
	"Chile/Continental",
	"Chile/EasterIsland",
	"Cuba",
	"EET",
	"EST",
	"EST5EDT",
	"Egypt",
	"Eire",
	"Etc/GMT+6",
	"Etc/GMT-12",
	"Etc/Universal",
	"Etc/GMT+10",
	"Etc/GMT-13",
	"Etc/GMT+3",
	"Etc/UTC",
	"Etc/GMT+4",
	"Etc/GMT+5",
	"Etc/UCT",
	"Etc/GMT0",
	"Etc/Greenwich",
	"Etc/GMT",
	"Etc/GMT-14",
	"Etc/GMT+9",
	"Etc/GMT-2",
	"Etc/GMT-8",
	"Etc/GMT-1",
	"Etc/GMT-7",
	"Etc/GMT-6",
	"Etc/GMT+2",
	"Etc/GMT+11",
	"Etc/GMT+8",
	"Etc/GMT-9",
	"Etc/Zulu",
	"Etc/GMT+0",
	"Etc/GMT+7",
	"Etc/GMT-0",
	"Etc/GMT-10",
	"Etc/GMT-3",
	"Etc/GMT-11",
	"Etc/GMT-5",
	"Etc/GMT-4",
	"Etc/GMT+1",
	"Etc/GMT+12",
	"Europe/Belfast",
	"Europe/Tallinn",
	"Europe/Amsterdam",
	"Europe/London",
	"Europe/Budapest",
	"Europe/Bucharest",
	"Europe/Tiraspol",
	"Europe/Tirane",
	"Europe/Volgograd",
	"Europe/Minsk",
	"Europe/Kaliningrad",
	"Europe/Zaporozhye",
	"Europe/Athens",
	"Europe/Kiev",
	"Europe/Istanbul",
	"Europe/Rome",
	"Europe/Malta",
	"Europe/Busingen",
	"Europe/Oslo",
	"Europe/Copenhagen",
	"Europe/Sofia",
	"Europe/Simferopol",
	"Europe/Vilnius",
	"Europe/Luxembourg",
	"Europe/Vatican",
	"Europe/Mariehamn",
	"Europe/Brussels",
	"Europe/Andorra",
	"Europe/Paris",
	"Europe/Vienna",
	"Europe/Belgrade",
	"Europe/Prague",
	"Europe/Moscow",
	"Europe/Podgorica",
	"Europe/Sarajevo",
	"Europe/Samara",
	"Europe/Stockholm",
	"Europe/San_Marino",
	"Europe/Jersey",
	"Europe/Berlin",
	"Europe/Warsaw",
	"Europe/Riga",
	"Europe/Madrid",
	"Europe/Vaduz",
	"Europe/Zurich",
	"Europe/Dublin",
	"Europe/Uzhgorod",
	"Europe/Zagreb",
	"Europe/Monaco",
	"Europe/Gibraltar",
	"Europe/Bratislava",
	"Europe/Nicosia",
	"Europe/Helsinki",
	"Europe/Isle_of_Man",
	"Europe/Ljubljana",
	"Europe/Skopje",
	"Europe/Lisbon",
	"Europe/Chisinau",
	"Europe/Guernsey",
	"Factory",
	"GB",
	"GB-Eire",
	"GMT",
	"GMT+0",
	"GMT-0",
	"GMT0",
	"Greenwich",
	"HST",
	"Hongkong",
	"Iceland",
	"Indian/Mayotte",
	"Indian/Antananarivo",
	"Indian/Christmas",
	"Indian/Maldives",
	"Indian/Mauritius",
	"Indian/Comoro",
	"Indian/Cocos",
	"Indian/Kerguelen",
	"Indian/Chagos",
	"Indian/Mahe",
	"Indian/Reunion",
	"Iran",
	"Israel",
	"Jamaica",
	"Japan",
	"Kwajalein",
	"Libya",
	"MET",
	"MST",
	"MST7MDT",
	"Mexico/General",
	"Mexico/BajaNorte",
	"Mexico/BajaSur",
	"Mideast/Riyadh87",
	"Mideast/Riyadh88",
	"Mideast/Riyadh89",
	"NZ",
	"NZ-CHAT",
	"Navajo",
	"PRC",
	"PST8PDT",
	"Pacific/Midway",
	"Pacific/Pitcairn",
	"Pacific/Tahiti",
	"Pacific/Fiji",
	"Pacific/Wallis",
	"Pacific/Samoa",
	"Pacific/Funafuti",
	"Pacific/Apia",
	"Pacific/Port_Moresby",
	"Pacific/Honolulu",
	"Pacific/Enderbury",
	"Pacific/Wake",
	"Pacific/Fakaofo",
	"Pacific/Pago_Pago",
	"Pacific/Kwajalein",
	"Pacific/Truk",
	"Pacific/Yap",
	"Pacific/Auckland",
	"Pacific/Johnston",
	"Pacific/Saipan",
	"Pacific/Tarawa",
	"Pacific/Chuuk",
	"Pacific/Kiritimati",
	"Pacific/Galapagos",
	"Pacific/Easter",
	"Pacific/Chatham",
	"Pacific/Pohnpei",
	"Pacific/Tongatapu",
	"Pacific/Nauru",
	"Pacific/Majuro",
	"Pacific/Guam",
	"Pacific/Guadalcanal",
	"Pacific/Marquesas",
	"Pacific/Gambier",
	"Pacific/Rarotonga",
	"Pacific/Noumea",
	"Pacific/Ponape",
	"Pacific/Kosrae",
	"Pacific/Palau",
	"Pacific/Niue",
	"Pacific/Norfolk",
	"Pacific/Efate",
	"Poland",
	"Portugal",
	"ROC",
	"ROK",
	"Singapore",
	"Turkey",
	"UCT",
	"US/Arizona",
	"US/Samoa",
	"US/Hawaii",
	"US/Central",
	"US/Michigan",
	"US/Alaska",
	"US/Aleutian",
	"US/Indiana-Starke",
	"US/Mountain",
	"US/Pacific",
	"US/Eastern",
	"US/Pacific-New",
	"US/East-Indiana",
	"UTC",
	"Universal",
	"W-SU",
	"WET",
	"Zulu",
}
